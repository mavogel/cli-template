# CLI Template

A comprehensive Go CLI application template with modern tooling and best practices.

## Overview

CLI Template provides a solid foundation for building command-line applications in Go, featuring:

- **Modern CLI Framework**: Built with [Cobra](https://cobra.dev/) for powerful command-line interfaces
- **Comprehensive Testing**: Unit tests with coverage reporting
- **Code Quality**: Integrated linting with golangci-lint
- **Automated Releases**: GoReleaser with GitHub Actions
- **Cross-Platform**: Builds for Linux, macOS, and Windows
- **Documentation**: MkDocs with Material theme

## Features

### 🚀 Development Experience
- Hot reloading during development
- Comprehensive Makefile for common tasks
- IDE-friendly project structure
- Git hooks for code quality

### 🔧 Build & Release
- Cross-platform compilation
- Automated semantic versioning
- GitHub Releases integration
- Homebrew tap support
- Docker container builds

### 📊 Code Quality
- Comprehensive linting rules
- Test coverage reporting
- Security scanning
- Dependency vulnerability checks

### 🔄 CI/CD
- GitHub Actions workflows
- Automated testing on multiple Go versions
- Release automation
- Documentation deployment

## Quick Start

```bash
# Clone the template
git clone https://github.com/mavogel/cli-template.git
cd cli-template

# Install dependencies
go mod download

# Build the application
make build

# Run tests
make test

# Run with linting
make lint

# Create a snapshot release
make release-snapshot
```

## Project Structure

```
cli-template/
├── cmd/                    # Command implementations
│   ├── root.go            # Root command
│   ├── hello.go           # Example subcommand
│   └── *_test.go          # Command tests
├── docs/                  # Documentation source
├── .github/
│   └── workflows/         # GitHub Actions
├── .goreleaser.yaml       # Release configuration
├── .golangci.yml         # Linting configuration
├── mkdocs.yml            # Documentation configuration
├── Makefile              # Build automation
├── main.go               # Application entry point
└── go.mod                # Go module definition
```

## Next Steps

- [Installation Guide](getting-started/installation.md) - Set up your environment
- [Quick Start](getting-started/quick-start.md) - Build your first command
- [Development Guide](development/project-structure.md) - Understand the codebase
- [Deployment](deployment/releases.md) - Release your application