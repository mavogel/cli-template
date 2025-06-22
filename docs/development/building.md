# Building

Learn how to build your CLI application for development and production.

## Development Builds

### Quick Build

```bash
# Build for current platform
make build

# Build with verbose output
go build -v -o bin/cli-template ./main.go
```

### Development Mode

```bash
# Run without building
make run

# Run with arguments
go run main.go hello --name "Developer"

# Development with hot reload (manual)
make dev
```

## Production Builds

### Single Platform

```bash
# Build for current platform with optimizations
go build -ldflags "-s -w" -o bin/cli-template ./main.go

# Build with version information
go build -ldflags "-s -w -X main.version=v1.0.0 -X main.commit=$(git rev-parse HEAD) -X main.date=$(date -u +%Y-%m-%dT%H:%M:%SZ)" -o bin/cli-template ./main.go
```

### Cross-Platform Builds

```bash
# Create snapshot release for all platforms
make release-snapshot

# Build specific platform
GOOS=linux GOARCH=amd64 go build -o bin/cli-template-linux-amd64 ./main.go
GOOS=windows GOARCH=amd64 go build -o bin/cli-template-windows-amd64.exe ./main.go
GOOS=darwin GOARCH=amd64 go build -o bin/cli-template-darwin-amd64 ./main.go
```

## Build Optimization

### Binary Size

```bash
# Reduce binary size
go build -ldflags "-s -w" ./main.go

# Further optimization with UPX (if installed)
upx --brute bin/cli-template
```

### Build Flags

Common build flags and their purposes:

- **`-s`**: Strip symbol table
- **`-w`**: Strip debug information
- **`-X`**: Set string variable at build time
- **`-trimpath`**: Remove file system paths
- **`-buildmode`**: Set build mode

### CGO Control

```bash
# Disable CGO for pure Go builds
CGO_ENABLED=0 go build ./main.go

# Enable static linking
CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' ./main.go
```

## GoReleaser Integration

### Configuration Overview

The `.goreleaser.yaml` file controls automated builds:

```yaml
builds:
  - main: ./main.go
    binary: cli-template
    env:
      - CGO_ENABLED=0
    goos:
      - linux
      - windows  
      - darwin
    goarch:
      - amd64
      - arm64
    ldflags:
      - -s -w
      - -X main.version={{.Version}}
      - -X main.commit={{.Commit}}
      - -X main.date={{.Date}}
```

### Local Testing

```bash
# Validate configuration
make release-check

# Create snapshot build
make release-snapshot

# Inspect built artifacts
ls -la dist/
```

### Build Matrix

GoReleaser builds for multiple targets:

| OS      | Architecture | Output                              |
|---------|-------------|-------------------------------------|
| Linux   | amd64       | `cli-template_linux_amd64`        |
| Linux   | arm64       | `cli-template_linux_arm64`        |
| macOS   | amd64       | `cli-template_darwin_amd64`       |
| macOS   | arm64       | `cli-template_darwin_arm64`       |
| Windows | amd64       | `cli-template_windows_amd64.exe`  |

## Build Targets

### Makefile Targets

Available make targets:

```bash
make help           # Show all available targets
make build          # Build for current platform
make test           # Run tests
make lint           # Run linting
make clean          # Clean build artifacts
make deps           # Download dependencies
make all            # Run all checks and build
make release-check  # Validate GoReleaser config
make release-snapshot # Create snapshot build
```

### Custom Targets

Add custom build targets to Makefile:

```makefile
# Custom target example
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux $(MAIN_PATH)

build-windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-windows.exe $(MAIN_PATH)

build-all: build-linux build-windows build
```

## Environment Variables

### Build Environment

```bash
# Go build environment
export GOOS=linux          # Target operating system
export GOARCH=amd64        # Target architecture
export CGO_ENABLED=0       # Disable CGO
export GO111MODULE=on      # Enable Go modules

# Build optimization
export GOPROXY=https://proxy.golang.org,direct
export GOSUMDB=sum.golang.org
```

### Version Information

Inject version information at build time:

```bash
VERSION=$(git describe --tags --always --dirty)
COMMIT=$(git rev-parse HEAD)
DATE=$(date -u +%Y-%m-%dT%H:%M:%SZ)

go build -ldflags "-X main.version=$VERSION -X main.commit=$COMMIT -X main.date=$DATE" ./main.go
```

## Docker Builds

### Multi-stage Dockerfile

```dockerfile
# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cli-template ./main.go

# Final stage
FROM scratch
COPY --from=builder /app/cli-template /usr/local/bin/cli-template
ENTRYPOINT ["/usr/local/bin/cli-template"]
```

### Docker Build Commands

```bash
# Build Docker image
docker build -t cli-template:latest .

# Build with build arguments
docker build --build-arg VERSION=v1.0.0 -t cli-template:v1.0.0 .

# Multi-platform build
docker buildx build --platform linux/amd64,linux/arm64 -t cli-template:latest .
```

## Troubleshooting

### Common Issues

**Module not found**
```bash
go mod tidy
go mod download
```

**CGO errors on cross-compilation**
```bash
CGO_ENABLED=0 go build ./main.go
```

**Version information not embedded**
```bash
# Check ldflags syntax
go build -ldflags "-X main.version=test" ./main.go
./cli-template --version
```

### Debug Builds

```bash
# Build with debug information
go build -gcflags "-N -l" ./main.go

# Race detection build
go build -race ./main.go

# Verbose build output
go build -v -x ./main.go
```