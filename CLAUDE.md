# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a comprehensive Go CLI application template built with the Cobra framework. It provides modern tooling, automated testing, releases, and documentation suitable for production-ready command-line tools.

**Language**: Go 1.23.10
**Framework**: Cobra (CLI framework)
**Build Tool**: Make
**Release Tool**: GoReleaser v2
**Documentation**: MkDocs with Material theme
**CI/CD**: GitHub Actions

## Key Commands

```bash
# Development
make build              # Build binary to bin/cli-template
make test               # Run tests with race detection
make test-coverage      # Generate coverage report (coverage.html)
make lint               # Run golangci-lint
make lint-fix           # Auto-fix linting issues
make all                # Clean, download deps, lint, test, and build

# Running
make run                # Run the application directly
go run main.go hello    # Run specific command during development

# Documentation (requires Docker)
make docs-serve         # Serve at http://localhost:8000
make docs-build         # Build to site/ directory

# Release
make release-check      # Validate GoReleaser config
make release-snapshot   # Test release without publishing
make release            # Create release (requires git tag)

# Testing specific files
go test ./cmd -run TestHelloCommand
go test -v ./cmd/hello_test.go
```

## Architecture

The project follows standard Go CLI patterns with Cobra:

### Command Structure
- `main.go`: Entry point, initializes and executes root command
- `cmd/root.go`: Root command setup, version handling, config initialization
- `cmd/hello.go`: Example subcommand implementation
- Each command has a corresponding `*_test.go` file

### Adding New Commands
Create `cmd/newcommand.go`:
```go
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
    Use:   "new",
    Short: "Brief description",
    Long:  `Detailed description of the command.`,
    RunE: func(cmd *cobra.Command, args []string) error {
        // Implementation
        return nil
    },
}

func init() {
    rootCmd.AddCommand(newCmd)
    // Add flags
    newCmd.Flags().StringP("name", "n", "", "description")
}
```

### Testing Pattern
Tests use a table-driven approach:
```go
func TestNewCommand(t *testing.T) {
    tests := []struct {
        name    string
        args    []string
        wantErr bool
        output  string
    }{
        {
            name:   "valid input",
            args:   []string{"--name", "test"},
            output: "expected output",
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

## Development Guidelines

1. **Go Version**: Use Go 1.23.10 or later
2. **Dependencies**: Only add necessary dependencies, run `go mod tidy` after changes
3. **Error Handling**: Always check and handle errors appropriately
4. **Testing**: Maintain >80% code coverage, write tests alongside implementation
5. **Linting**: Code must pass all linters in `.golangci.yml`
6. **Documentation**: Update command help text and docs/ when adding features

## Commit Message Format

Use conventional commits:
- `feat:` New feature
- `fix:` Bug fix
- `docs:` Documentation changes
- `style:` Code style changes (formatting, etc.)
- `refactor:` Code refactoring
- `test:` Test additions or changes
- `chore:` Maintenance tasks

## CI/CD Pipeline

GitHub Actions workflows:
- **CI** (`ci.yml`): Runs on every push
  - Tests on Go 1.21.x, 1.22.x, 1.23.x
  - Linting with golangci-lint
  - Build verification
- **Release** (`release.yml`): Triggered by version tags
  - Cross-platform builds (Linux, macOS, Windows)
  - Docker image publishing
  - Homebrew tap updates
- **Docs** (`docs.yml`): Documentation deployment to GitHub Pages

## Configuration

The CLI supports configuration via:
1. `./cli-template.yaml` (current directory)
2. `~/.cli-template.yaml` (home directory)
3. `/etc/cli-template/config.yaml` (system-wide)

Viper automatically merges these in order of precedence.

## Important Notes

- Binary name is `cli-template` - update in Makefile and .goreleaser.yaml when forking
- Version is injected during build from git tags
- GoReleaser expects semantic version tags (v1.0.0)
- Docker builds use scratch image for minimal size
- Shell completions are auto-generated for bash/zsh
- All commands should have comprehensive help text
- Use `cobra.ExactArgs()`, `cobra.MinimumNArgs()` for argument validation
- Prefer `RunE` over `Run` for proper error handling in commands