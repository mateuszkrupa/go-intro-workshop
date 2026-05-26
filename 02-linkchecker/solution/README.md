# Link Checker API — Solution

Complete reference implementation of the concurrent link checker HTTP API. This solution demonstrates goroutines, channels, `sync.WaitGroup`, and mutex-protected shared state in Go.

The server exposes two endpoints:

- **POST /check** — Accepts a JSON body with a `urls` array (1–20 entries), checks each URL concurrently using one goroutine per URL, and returns status codes with response times in milliseconds.
- **GET /history** — Returns the last 10 check operations, ordered most recent first.

## Run

```bash
# Start the server (default port 8080)
go run .
```

To use a custom port:

**macOS/Linux:**
```bash
PORT=9090 go run .
```

**Windows (PowerShell):**
```powershell
$env:PORT="9090"; go run .
```

## Test

```bash
# Compile the project
go build ./...

# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

**Manual testing with curl:**

**macOS/Linux:**
```bash
# Check URLs
curl -X POST http://localhost:8080/check \
  -H "Content-Type: application/json" \
  -d '{"urls": ["https://go.dev", "https://github.com", "https://example.com"]}'

# View history
curl http://localhost:8080/history
```

**Windows (PowerShell):**
```powershell
# Check URLs
Invoke-RestMethod -Method POST -Uri http://localhost:8080/check `
  -ContentType "application/json" `
  -Body '{"urls": ["https://go.dev", "https://github.com", "https://example.com"]}'

# View history
Invoke-RestMethod http://localhost:8080/history
```
