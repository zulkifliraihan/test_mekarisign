# Hot Reload Development Guide

## Overview

This application uses **Air** for hot reload development, similar to `nodemon` or `npm run dev` in Node.js. With Air, the server automatically restarts whenever code changes.

## Installation

```bash
go install github.com/air-verse/air@latest
```

## Usage

### Quick Start

```bash
# Using Makefile (Recommended)
make dev

# Or run Air directly
air
```

### What Happens?

When you run `make dev` or `air`:

1. ✅ Air builds the app to `tmp/main`
2. ✅ The server starts automatically
3. ✅ Air watches all `.go` files
4. ✅ On file changes, Air will:
   - Stop the running server
   - Rebuild the app
   - Restart the server with new code

## Configuration

Air configuration is in `.air.toml`:

```toml
[build]
  cmd = "go build -o ./tmp/main ./cmd/api"  # Build command
  bin = "./tmp/main"                         # Binary location
  include_ext = ["go", "tpl", "tmpl", "html"] # File extensions to watch
  exclude_dir = ["assets", "tmp", "vendor"]  # Directories to ignore
  delay = 1000                               # Delay before rebuild (ms)
```

## Features

### Auto Restart
```bash
make dev

# Now edit any file in internal/*
# The server will restart automatically!
```

### Colored Output
Air provides colored output for easy identification:
- 🟢 **Green** - Runner messages
- 🟡 **Yellow** - Build messages
- 🔴 **Magenta** - Main app output
- 🔵 **Cyan** - Watcher messages

### Build Error Logs
If a build error occurs, Air saves the error log to `build-errors.log`

```bash
# Check build errors
cat build-errors.log
```

## Directory Structure

```
test_mekari/
├── .air.toml              # Air configuration
├── tmp/                   # Temporary build directory (gitignored)
│   └── main              # Built binary
└── build-errors.log      # Build error logs (gitignored)
```

## Commands

| Command | Description |
|---------|-------------|
| `make dev` | Start with hot reload |
| `make run` | Start without hot reload |
| `make clean` | Clean build artifacts (includes tmp/) |

## Comparison with Node.js

### Node.js (npm run dev)
```bash
npm run dev  # Auto-restart with nodemon
```

### Go (make dev)
```bash
make dev     # Auto-restart with Air
```

**Both provide:**
- ✅ Auto-restart on file changes
- ✅ Fast rebuild
- ✅ Development-friendly output
- ✅ Error logging

## Tips

### 1. Exclude Files from Watch

Edit `.air.toml` to exclude files:

```toml
[build]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
```

### 2. Custom Delay

If rebuilds are too fast, adjust the delay:

```toml
[build]
  delay = 1000  # ms before rebuild
```

### 3. Clean Before Start

If issues occur, clean first:

```bash
make clean
make dev
```

## Troubleshooting

### Issue: "air: command not found"

**Solution:**
```bash
# Install Air
go install github.com/air-verse/air@latest

# Make sure $GOPATH/bin is in your PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

### Issue: Server does not restart automatically

**Solution:**
```bash
# 1. Stop Air (Ctrl+C)
# 2. Clean tmp directory
make clean
# 3. Restart Air
make dev
```

### Issue: Port already in use

**Solution:**
```bash
# Kill process on port 8080
lsof -ti:8080 | xargs kill -9

# Or use custom port
PORT=3000 air
```

### Issue: Build errors not showing

**Solution:**
```bash
# Check build error log
cat build-errors.log

# Or run with verbose
air -d  # debug mode
```

## Advanced Configuration

### Watch Specific Directories Only

```toml
[build]
  include_dir = ["cmd", "internal"]
```

### Custom Build Command

```toml
[build]
  cmd = "go build -tags dev -o ./tmp/main ./cmd/api"
```

### Pre/Post Commands

```toml
[build]
  pre_cmd = ["swag init"]              # Run before build
  post_cmd = ["echo 'Build complete'"] # Run after build
```

## Performance

### Build Time
- First build: ~2-3 seconds
- Rebuild on change: ~1-2 seconds

### Memory Usage
- Air process: ~20-30 MB
- Your app: depends on your code

## Best Practices

1. **Use make dev for development**
   ```bash
   make dev  # Not go run
   ```

2. **Use make build for production**
   ```bash
   make build
   ./todo-api
   ```

3. **Clean regularly**
   ```bash
   make clean  # Before commit
   ```

4. **Commit .air.toml**
   - Share configuration with team
   - Keep it in version control

5. **Ignore tmp/ and build-errors.log**
   - Already in .gitignore
   - Don't commit build artifacts

## Resources

- **Air Repository**: https://github.com/air-verse/air
- **Air Documentation**: https://github.com/air-verse/air/blob/master/README.md
- **Configuration Examples**: https://github.com/air-verse/air/blob/master/air_example.toml

## Comparison Table

| Feature | `go run` | `make run` | `make dev` (Air) |
|---------|----------|------------|------------------|
| Auto-restart | ❌ | ❌ | ✅ |
| Fast rebuild | ❌ | ❌ | ✅ |
| Colored output | ❌ | ❌ | ✅ |
| Error logging | ❌ | ❌ | ✅ |
| File watching | ❌ | ❌ | ✅ |
| Production ready | ❌ | ❌ | ❌ |

For production, always use:
```bash
make build
./todo-api
```

---

**Happy Coding with Hot Reload!** 🔥
