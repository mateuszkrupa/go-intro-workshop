# Go in 90 minutes: from zero to concurrent API

A hands-on workshop. In 90 minutes you'll go from writing your first Go function to building a concurrent HTTP API — no prior Go knowledge required.

## Quick Navigation

- [Pre-workshop setup instructions](PREWORK.md)
- [Project 1: Calculator (starter code)](01-calculator/starter/)
- [Project 2: Link Checker API (starter code)](02-linkchecker/starter/)
- [Slide deck](slides/)

## Troubleshooting

### Go installation verification failure

**Symptom:** Running `go version` returns "command not found" or shows a version below 1.26.

**Fix:**
1. Re-download Go from https://go.dev/dl and follow the install instructions for your OS.
2. Ensure the Go binary is on your `PATH`:
   - macOS/Linux: add `export PATH=$PATH:/usr/local/go/bin` to your shell profile (`~/.bashrc`, `~/.zshrc`).
   - Windows: the installer should set this automatically — restart your terminal after installing.
3. Open a **new** terminal window and run `go version` again.

### IDE / gopls issues

**Symptom:** VS Code shows "gopls not found" errors, red squiggles everywhere, or autocomplete doesn't work.

**Fix:**
1. Install gopls manually: `go install golang.org/x/tools/gopls@latest`
2. Make sure `$(go env GOPATH)/bin` is on your `PATH`.
3. In VS Code, open the Command Palette (Ctrl+Shift+P / Cmd+Shift+P) and run **Go: Restart Language Server**.
4. If problems persist, try removing the Go extension and reinstalling it.

### Module download errors

**Symptom:** `go mod tidy` or `go mod download` fails with network timeouts or checksum mismatches.

**Fix:**
1. Check your internet connection — module downloads require network access.
2. If you're behind a corporate proxy, set the `HTTPS_PROXY` environment variable.
3. Try setting a mirror: `go env -w GOPROXY=https://proxy.golang.org,direct`
4. If all else fails, ask the instructor for a USB stick with cached modules, or use the [Go Playground](https://go.dev/play) as a fallback.
