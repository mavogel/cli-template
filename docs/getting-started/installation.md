# Installation

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.21+**: [Download Go](https://golang.org/dl/)
- **Git**: [Download Git](https://git-scm.com/downloads)
- **Make**: Usually comes with your OS or development tools

## Installation Methods

### Option 1: Go Install (Recommended)

```bash
go install github.com/mavogel/cli-template@latest
```

### Option 2: Download Binary

1. Visit the [Releases page](https://github.com/mavogel/cli-template/releases)
2. Download the appropriate binary for your OS
3. Extract and place in your PATH

### Option 3: Homebrew (macOS/Linux)

```bash
brew tap mavogel/homebrew-tap
brew install cli-template
```

### Option 4: Build from Source

```bash
# Clone the repository
git clone https://github.com/mavogel/cli-template.git
cd cli-template

# Build the binary
make build

# Install to your PATH
make install
```

## Verify Installation

```bash
cli-template --version
cli-template --help
```

## Development Setup

If you plan to contribute or modify the code:

```bash
# Clone the repository
git clone https://github.com/mavogel/cli-template.git
cd cli-template

# Install dependencies
go mod download

# Install development tools
make dev-tools

# Run tests
make test

# Run linting
make lint
```

### Development Dependencies

The following tools are recommended for development:

- **golangci-lint**: Code linting
- **GoReleaser**: Release automation
- **MkDocs**: Documentation generation

Install them with:

```bash
# golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# GoReleaser
go install github.com/goreleaser/goreleaser@latest

# Docker (for MkDocs documentation)
brew install --cask docker
```