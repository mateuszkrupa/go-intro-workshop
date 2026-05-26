# Project 2: Concurrent Link Checker API

Build a concurrent HTTP API in Go that checks the status of multiple URLs in parallel.

## How to Build & Run

```bash
# Start the server (port 8080)
go run .

# Verify it compiles
go build ./...
```

> **Windows note:** All `go` commands work the same in PowerShell and Command Prompt. Use PowerShell for the best experience. For `curl` commands below, see the Windows alternatives provided.

**Test with curl (macOS/Linux):**

```bash
# Check multiple URLs
curl -X POST http://localhost:8080/check \
  -H "Content-Type: application/json" \
  -d '{"urls": ["https://go.dev", "https://github.com"]}'

# View check history
curl http://localhost:8080/history
```

**Test with PowerShell (Windows):**

```powershell
# Check multiple URLs
Invoke-RestMethod -Method POST -Uri http://localhost:8080/check `
  -ContentType "application/json" `
  -Body '{"urls": ["https://go.dev", "https://github.com"]}'

# View check history
Invoke-RestMethod http://localhost:8080/history
```

---

## Phase 1: Store (mutex-protected history)

Implement `Add` and `GetHistory` on `Store` in `store.go`.

### Tasks

1. Implement `Add(response CheckResponse)`:
   - Lock `s.mu` before writing, unlock after
   - Prepend the new response: `s.history = append([]CheckResponse{response}, s.history...)`
   - Cap at 10 entries: `if len(s.history) > 10 { s.history = s.history[:10] }`
2. Implement `GetHistory() []CheckResponse`:
   - Lock `s.mu` before reading, unlock after
   - Return a copy: `result := make([]CheckResponse, len(s.history))` + `copy(result, s.history)`
3. Start the server and verify it starts without errors:
   ```bash
   go run .
   ```

### Hints

- `s.mu.Lock()` blocks other goroutines until `s.mu.Unlock()` is called — safe for concurrent access.
- Use `defer s.mu.Unlock()` right after `s.mu.Lock()` so you never forget to unlock.
- A copy protects the caller from accidentally mutating internal state.

---

## Phase 2: POST /check (sequential)

Wire up the handler in `handleCheck` — JSON in, URL checks, JSON out.

### Tasks

1. Read through the top of `handleCheck` — JSON decoding and validation are already done for you.
2. Find the `// TODO` comment about looping over URLs. Write a for-loop that checks each URL sequentially.
3. Remove the `_ = results` line (it was just there to suppress the compiler error).
4. Below the loop, replace the placeholder lines (`_ = store`, and the `writeError` call) with code that:
   - Wraps `results` in a `CheckResponse`
   - Saves it to the store
   - Sets the `Content-Type` header to `application/json`
   - Encodes the response as JSON to `w`
5. Test:

   **macOS/Linux:**
   ```bash
   curl -X POST http://localhost:8080/check \
     -H "Content-Type: application/json" \
     -d '{"urls": ["https://go.dev", "https://example.com"]}'
   ```

   **Windows (PowerShell):**
   ```powershell
   Invoke-RestMethod -Method POST -Uri http://localhost:8080/check `
     -ContentType "application/json" `
     -Body '{"urls": ["https://go.dev", "https://example.com"]}'
   ```

   You should see results — but it will be slow (sequential).

### Hints

- `checkURL` is already implemented in `checker.go` — just call it.
- `writeError` is already implemented — use it for bad input.
- `json.NewEncoder(w).Encode(v)` streams JSON directly to the response writer.

---

## Phase 3: Make it concurrent

Replace the sequential for-loop with goroutines, a channel, and `sync.WaitGroup`.

### Tasks

1. Add `"sync"` to your imports.
2. Replace the sequential for-loop with:
   ```go
   ch := make(chan CheckResult, len(req.URLs))
   var wg sync.WaitGroup

   for _, url := range req.URLs {
       wg.Add(1)
       go func(u string) {
           defer wg.Done()
           ch <- checkURL(u)
       }(url)
   }

   go func() {
       wg.Wait()
       close(ch)
   }()

   var results []CheckResult
   for result := range ch {
       results = append(results, result)
   }
   ```
3. Test again with the same curl — it should be noticeably faster for multiple URLs.
4. Try with 5 URLs and compare timing:

   **macOS/Linux:**
   ```bash
   curl -X POST http://localhost:8080/check \
     -H "Content-Type: application/json" \
     -d '{"urls":["https://go.dev","https://github.com","https://example.com","https://cloudflare.com","https://pkg.go.dev"]}'
   ```

   **Windows (PowerShell):**
   ```powershell
   Invoke-RestMethod -Method POST -Uri http://localhost:8080/check `
     -ContentType "application/json" `
     -Body '{"urls":["https://go.dev","https://github.com","https://example.com","https://cloudflare.com","https://pkg.go.dev"]}'
   ```

### Hints

- Buffered channel `make(chan CheckResult, len(req.URLs))` prevents goroutines from blocking each other.
- `defer wg.Done()` inside the goroutine ensures Done is always called, even on panic.
- Closing the channel after `wg.Wait()` lets `for result := range ch` know when to stop.

---

## Phase 4: GET /history

Implement `handleHistory` in `main.go`.

### Tasks

1. Replace the stub in `handleHistory` with code that:
   - Gets the history from the store
   - Sets the `Content-Type` header to `application/json`
   - Encodes the history as JSON to `w`
2. Remove the `_ = store` line.
3. Test:

   **macOS/Linux:**
   ```bash
   # Run a few checks first, then:
   curl http://localhost:8080/history
   ```

   **Windows (PowerShell):**
   ```powershell
   # Run a few checks first, then:
   Invoke-RestMethod http://localhost:8080/history
   ```

   You should see the last checks ordered newest-first.

---

## Run & Test Commands

| Command | Purpose |
|---------|---------|
| `go run .` | Start the server on port 8080 |
| `go build ./...` | Verify the project compiles |

## Stretch Goals

1. **Rate limiter** — limit concurrent requests with a buffered semaphore channel: `sem := make(chan struct{}, 5)`
2. **Persist to file** — save history to `data.json` on each check, load on startup
3. **Graceful shutdown** — use `signal.Notify` + `http.Server.Shutdown` to drain in-flight requests on Ctrl+C
