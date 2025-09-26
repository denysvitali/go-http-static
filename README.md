# Go HTTP Static

A fast, simple HTTP server for serving static files with optional TLS support.

## Features

- 🚀 **Fast** - Built with Go's standard library for maximum performance
- 🔒 **TLS Support** - Optional HTTPS with custom certificates
- 🎨 **Styled Output** - Beautiful terminal interface with colors
- 📝 **Structured Logging** - Comprehensive logging with Logrus
- 🔧 **Modern CLI** - Built with Cobra for excellent UX
- 📦 **Static Binaries** - No dependencies, runs anywhere

## Installation

### Download Pre-built Binaries

Download the latest release for your platform:
- [Linux AMD64](https://github.com/denysvitali/go-http-static/releases/latest/download/go-http-static-linux-amd64)
- [Linux ARM64](https://github.com/denysvitali/go-http-static/releases/latest/download/go-http-static-linux-arm64)

### Build from Source

```bash
git clone https://github.com/denysvitali/go-http-static.git
cd go-http-static
go build -o go-http-static ./cmd
```

## Usage

### Basic HTTP Server

```bash
# Serve current directory on port 8080
./go-http-static serve .

# Serve specific directory on custom port
./go-http-static serve /path/to/files --port 9000

# Bind to specific interface
./go-http-static serve /var/www --listen 127.0.0.1 --port 8080
```

### HTTPS Server

```bash
# Serve with TLS
./go-http-static serve /var/www --tls --cert server.crt --key server.key --port 443
```

### Command Help

```bash
# View all commands
./go-http-static --help

# View serve command options
./go-http-static serve --help

# Check version
./go-http-static --version
```

## Command Line Options

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--port` | `-p` | Port to listen on | `8080` |
| `--listen` | `-l` | Listen address | All interfaces |
| `--tls` | `-t` | Enable HTTPS/TLS | `false` |
| `--cert` | `-c` | TLS certificate file | |
| `--key` | `-k` | TLS private key file | |

## Examples

```bash
# Development server
./go-http-static serve ./public

# Production HTTPS server
./go-http-static serve /var/www/html \
  --port 443 \
  --tls \
  --cert /etc/ssl/certs/server.crt \
  --key /etc/ssl/private/server.key

# Local testing with custom port
./go-http-static serve ./dist --port 3000 --listen localhost
```

## Development

### Requirements

- Go 1.24+

### Building

```bash
# Build for current platform
go build -o go-http-static ./cmd

# Build static binary for Linux AMD64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-http-static-linux-amd64 ./cmd

# Build static binary for Linux ARM64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o go-http-static-linux-arm64 ./cmd
```

### Testing

```bash
go test ./...
```

## License

MIT License - see LICENSE file for details.