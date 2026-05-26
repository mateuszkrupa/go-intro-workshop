# Pre-Workshop Setup

Complete these steps **before** the workshop so you arrive with a working Go environment.

> **WiFi note:** The venue has WiFi, which you'll need for `go mod tidy` during the workshop. However, completing the offline prep below ensures you can work even if the network is unreliable.

---

## 1. Install Go 1.26+

Download the installer for your OS from **https://go.dev/dl**

### macOS

```bash
# Option A: Download the .pkg installer from https://go.dev/dl and run it
# Option B: Homebrew
brew install go
```

### Windows

1. Download the `.msi` installer from https://go.dev/dl
2. Run the installer (accepts defaults)
3. Open a **new** terminal (Command Prompt or PowerShell) after installation

### Linux

```bash
# Download the tarball (replace version as needed)
wget https://go.dev/dl/go1.26.2.linux-amd64.tar.gz

# Remove any previous installation and extract
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.26.2.linux-amd64.tar.gz

# Add to PATH (add this line to ~/.bashrc or ~/.zshrc)
export PATH=$PATH:/usr/local/go/bin
```

### Verify installation

Open a terminal and run:

```bash
go version
```

Expected output (version must be **1.26 or higher**):

```
go version go1.26.x <os>/<arch>
```

For example: `go version go1.26.2 darwin/arm64`

---

## 2. IDE Setup

Choose one of:

- **VS Code** (free) — install the [official Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- **GoLand** (JetBrains) — free educational license available

### Install gopls (Go language server)

```bash
go install golang.org/x/tools/gopls@latest
```

This gives you autocomplete, inline errors, and go-to-definition in your editor.

---

## 3. Smoke Test

Clone the workshop repository and verify the starter code compiles:

```bash
git clone <repository-url>
cd go-introduction/01-calculator/starter
go test ./...
```

**Expected result:** The code compiles successfully, but tests **fail** (non-zero exit code). This is correct — you'll make them pass during the workshop.

---

## 4. Offline Prep

To cache all dependencies locally (so you can work without WiFi):

```bash
# From the repository root
cd 01-calculator/starter && go mod download
cd ../../02-linkchecker/starter && go mod download
```

Verify the cache works by building with network disabled:

```bash
cd 01-calculator/starter && go build ./...
cd ../../02-linkchecker/starter && go build ./...
```

If both commands succeed without errors, you're ready to work offline.

---

## 5. Troubleshooting

| Problem | Solution |
|---------|----------|
| `go version` not found | Restart your terminal. On Linux, ensure `/usr/local/go/bin` is in your `PATH`. |
| gopls errors in VS Code | Run `Go: Install/Update Tools` from the command palette (Ctrl+Shift+P / Cmd+Shift+P). |
| `go mod download` fails | Check your internet connection. Try setting `GOPROXY=https://proxy.golang.org,direct`. |
| Nothing works locally | Use the **Go Playground** as a fallback: https://go.dev/play — it supports running and sharing Go code in the browser. |

If you're stuck, don't worry — we'll have time at the start of the workshop to help with setup issues.
