# CLI Template

A comprehensive Go CLI application template with modern tooling, automated testing, releases, and documentation.

[![CI](https://github.com/mavogel/cli-template/actions/workflows/ci.yml/badge.svg)](https://github.com/mavogel/cli-template/actions/workflows/ci.yml)
[![Release](https://github.com/mavogel/cli-template/actions/workflows/release.yml/badge.svg)](https://github.com/mavogel/cli-template/actions/workflows/release.yml)
[![Documentation](https://github.com/mavogel/cli-template/actions/workflows/docs.yml/badge.svg)](https://github.com/mavogel/cli-template/actions/workflows/docs.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mavogel/cli-template)](https://goreportcard.com/report/github.com/mavogel/cli-template)

## 🚀 Features

- **Modern CLI Framework**: Built with [Cobra](https://cobra.dev/)
- **Comprehensive Testing**: Unit tests with coverage reporting
- **Code Quality**: Integrated linting with golangci-lint
- **Automated Releases**: Cross-platform builds with GoReleaser
- **GitHub Actions**: CI/CD pipelines for testing, building, and deploying
- **Documentation**: MkDocs with Material theme, auto-deployed to GitHub Pages
- **Cross-Platform**: Builds for Linux, macOS, and Windows (amd64/arm64)
- **Container Support**: Docker images and Kubernetes deployment examples
- **Package Distribution**: Homebrew tap and GitHub Container Registry

## 📋 Table of Contents

- [Quick Start](#-quick-start)
- [Installation](#-installation)
- [Development Setup](#-development-setup)
- [Usage](#-usage)
- [Project Structure](#-project-structure)
- [Development](#-development)
- [Contributing](#-contributing)
- [License](#-license)

## ⚡ Quick Start

```bash
# Install the CLI
go install github.com/mavogel/cli-template@latest

# Try it out
cli-template hello
cli-template hello --name "Developer"
cli-template --version
```

## 📦 Installation

### Option 1: Go Install (Recommended)
```bash
go install github.com/mavogel/cli-template@latest
```

### Option 2: Homebrew
```bash
brew tap mavogel/homebrew-tap
brew install cli-template
```

### Option 3: Download Binary
Download the latest release from [GitHub Releases](https://github.com/mavogel/cli-template/releases).

### Option 4: Container
```bash
docker run --rm ghcr.io/mavogel/cli-template:latest --help
```

## 🛠 Development Setup

### Prerequisites

Install the required tools using Homebrew:

```bash
# Install Go (if not already installed)
brew install go

# Install development tools
brew install golangci-lint
brew install goreleaser

# Install Docker for documentation (MkDocs)
# Download from https://docker.com/products/docker-desktop
# Or install via Homebrew Cask:
brew install --cask docker
```

### Setup Project

```bash
# Clone the repository
git clone https://github.com/mavogel/cli-template.git
cd cli-template

# Install Go dependencies
go mod download

# Verify setup
make all

# Run tests
make test

# Build the application
make build

# Serve documentation locally
make docs-serve
```

### Development Tools Overview

| Tool | Purpose | Installation |
|------|---------|-------------|
| **Go** | Programming language | `brew install go` |
| **golangci-lint** | Code linting | `brew install golangci-lint` |
| **GoReleaser** | Release automation | `brew install goreleaser` |
| **Docker** | Documentation (MkDocs) | `brew install --cask docker` |
| **Make** | Build automation | Usually pre-installed |

## 🎯 Usage

### Basic Commands

```bash
# Show help
cli-template --help

# Show version
cli-template --version

# Hello command
cli-template hello
cli-template hello --name "Alice"
cli-template hello -n "Bob"

# Generate shell completions
cli-template completion bash > /usr/local/etc/bash_completion.d/cli-template
cli-template completion zsh > "${fpath[1]}/_cli-template"
```

### Configuration

The CLI looks for configuration files in:
- `./cli-template.yaml` (current directory)
- `~/.cli-template.yaml` (home directory)
- `/etc/cli-template/config.yaml` (system)

Example configuration:
```yaml
log_level: info
output_format: text
hello:
  default_name: "World"
```

## 📁 Project Structure

```
cli-template/
├── .github/
│   └── workflows/          # GitHub Actions CI/CD
├── cmd/                    # Command implementations
│   ├── root.go            # Root command
│   ├── hello.go           # Example subcommand
│   └── *_test.go          # Command tests
├── docs/                  # Documentation source
├── dist/                  # Build artifacts (generated)
├── bin/                   # Local build output
├── .golangci.yml         # Linting configuration
├── .goreleaser.yaml      # Release configuration
├── Dockerfile            # Container build
├── Makefile              # Build automation
├── mkdocs.yml            # Documentation config
├── go.mod                # Go module
└── main.go               # Application entry point
```

## 💻 Development

### Available Make Targets

```bash
make help               # Show all available targets
make build              # Build the binary
make test               # Run tests
make test-coverage      # Run tests with coverage
make lint               # Run linting
make lint-fix           # Run linting with auto-fix
make clean              # Clean build artifacts
make deps               # Download dependencies
make run                # Run the application
make all                # Run all checks and build

# Release
make release-check      # Validate GoReleaser config
make release-snapshot   # Create snapshot build
make release            # Create release (requires tag)

# Documentation (requires Docker)
make docs-serve         # Serve docs locally at http://localhost:8000
make docs-build         # Build documentation
make docs               # Build and validate docs
```

### Testing

```bash
# Run all tests
make test

# Run with coverage
make test-coverage
open coverage.html

# Run specific tests
go test ./cmd -run TestHelloCommand

# Benchmark tests
go test -bench=. ./...
```

### Linting

```bash
# Run all linters
make lint

# Auto-fix issues where possible
make lint-fix

# Run specific linter
golangci-lint run --enable-only=errcheck
```

### Building

```bash
# Local build
make build

# Cross-platform snapshot
make release-snapshot

# Check built artifacts
ls -la dist/
```

### Documentation

```bash
# Serve locally using Docker (http://localhost:8000)
make docs-serve

# Build documentation using Docker
make docs-build

# View built docs
open site/index.html
```

## 🚀 Releases

### Automated Releases

1. Create and push a tag:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

2. GitHub Actions automatically:
   - Builds cross-platform binaries
   - Creates GitHub release
   - Publishes Docker images
   - Updates Homebrew tap

### Manual Release Testing

```bash
# Test release configuration
make release-check

# Create snapshot (no publishing)
make release-snapshot

# Check artifacts
ls -la dist/
```

## 🔧 Customization

### Adding New Commands

1. Create `cmd/newcommand.go`:
```go
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
    Use:   "new",
    Short: "Description of new command",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("New command executed!")
    },
}

func init() {
    rootCmd.AddCommand(newCmd)
}
```

2. Add tests in `cmd/newcommand_test.go`
3. Update documentation

### Modifying for Your Project

1. Update `go.mod` with your module name
2. Replace `mavogel/cli-template` in all files
3. Update `.goreleaser.yaml` repository settings
4. Modify `mkdocs.yml` site information
5. Update GitHub Actions repository references

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guidelines](https://mavogel.github.io/cli-template/contributing/guidelines/) for details.

### Quick Contribution Steps

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Make your changes and add tests
4. Run tests and linting: `make all`
5. Commit with conventional commits: `git commit -m "feat: add amazing feature"`
6. Push and create a Pull Request

## 📖 Documentation

Full documentation is available at [https://mavogel.github.io/cli-template/](https://mavogel.github.io/cli-template/)

- [Installation Guide](https://mavogel.github.io/cli-template/getting-started/installation/)
- [Development Guide](https://mavogel.github.io/cli-template/development/project-structure/)
- [API Reference](https://mavogel.github.io/cli-template/reference/commands/)
- [Contributing](https://mavogel.github.io/cli-template/contributing/guidelines/)

## 🐛 Issues and Support

- **Bug Reports**: [GitHub Issues](https://github.com/mavogel/cli-template/issues)
- **Feature Requests**: [GitHub Issues](https://github.com/mavogel/cli-template/issues)
- **Questions**: [GitHub Discussions](https://github.com/mavogel/cli-template/discussions)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Cobra](https://cobra.dev/) - Powerful CLI framework
- [GoReleaser](https://goreleaser.com/) - Release automation
- [golangci-lint](https://golangci-lint.run/) - Go linting
- [MkDocs Material](https://squidfunk.github.io/mkdocs-material/) - Documentation theme

---

**Made with ❤️ for the Go community**