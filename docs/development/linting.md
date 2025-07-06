# Linting

Code quality enforcement and linting configuration for your CLI application.

## Overview

The project uses [golangci-lint](https://golangci-lint.run/) for comprehensive code quality checks. This ensures consistent code style, catches potential bugs, and enforces security best practices.

## Configuration

### `.golangci.yml`

The linting configuration is defined in `.golangci.yml`:

```yaml
run:
  timeout: 5m
  modules-download-mode: readonly

linters:
  enable:
    - errcheck      # Check for unchecked errors
    - gosimple      # Simplify code suggestions
    - govet         # Go vet analysis
    - ineffassign   # Detect ineffectual assignments
    - staticcheck   # Static analysis checks
    - typecheck     # Type checking
    - unused        # Find unused code
    - gofmt         # Format checking
    - goimports     # Import organization
    - misspell      # Spell checking
    - revive        # Golint replacement
    - gosec         # Security analysis
    - gocritic      # Comprehensive checks
    - gocyclo       # Cyclomatic complexity
    - goconst       # Repeated strings
    - godot         # Comment formatting
    - nolintlint    # Nolint directive checking
```

## Running Lints

### Basic Commands

```bash
# Run all linters
make lint

# Run with auto-fix where possible
make lint-fix

# Run specific linter
golangci-lint run --enable-only=errcheck

# Run on specific files/directories
golangci-lint run ./cmd/
```

### Detailed Output

```bash
# Verbose output
golangci-lint run -v

# Show all issues (not just new ones)
golangci-lint run --new=false

# Output in different formats
golangci-lint run --out-format=json
golangci-lint run --out-format=checkstyle
```

## Enabled Linters

### Error Handling

**errcheck**: Ensures all errors are properly handled
```go
// Bad
file, _ := os.Open("file.txt")

// Good  
file, err := os.Open("file.txt")
if err != nil {
    return err
}
```

**ineffassign**: Detects values assigned but never used
```go
// Bad
func example() {
    x := 10
    x = 20  // Previous assignment unused
    fmt.Println(x)
}
```

### Code Quality

**gosimple**: Suggests code simplifications
```go
// Bad
if len(slice) == 0 {
    return true
}
return false

// Good
return len(slice) == 0
```

**unused**: Finds unused variables, functions, types
```go
// Bad - unused variable
func example() {
    unused := "value"  // Will be flagged
    fmt.Println("hello")
}
```

### Formatting

**gofmt**: Enforces standard Go formatting
**goimports**: Organizes imports properly

```go
// Good imports organization
import (
    "fmt"           // Standard library
    "os"
    
    "github.com/spf13/cobra"  // Third-party
    
    "github.com/yourorg/yourapp/internal"  // Local
)
```

### Security

**gosec**: Security vulnerability detection
```go
// Bad - potential security issue
password := "hardcoded_password"  // G101: hardcoded credentials

// Bad - weak crypto
md5Hash := md5.Sum(data)  // G401: weak crypto
```

### Complexity

**gocyclo**: Cyclomatic complexity checking
```go
// Complex function - will be flagged if too many branches
func complexFunction(x int) string {
    if x > 10 {
        if x > 20 {
            if x > 30 {
                return "very high"
            }
            return "high"
        }
        return "medium"
    }
    return "low"
}
```

## Linter Settings

### Custom Configuration

```yaml
linters-settings:
  revive:
    min-confidence: 0
  gocyclo:
    min-complexity: 15
  goconst:
    min-len: 2
    min-occurrences: 2
  misspell:
    locale: US
  godot:
    capital: true
```

### Issue Exclusions

```yaml
issues:
  exclude-rules:
    - path: _test\.go
      linters:
        - gosec      # Allow security issues in tests
        - goconst    # Allow string duplication in tests
    - path: main\.go
      linters:
        - gochecknoinits  # Allow init functions in main
```

## Integration

### Editor Integration

#### VS Code

Install the Go extension and configure:

```json
{
    "go.lintTool": "golangci-lint",
    "go.lintFlags": ["--fast"],
    "go.lintOnSave": "package"
}
```

#### Vim/Neovim

With vim-go:
```vim
let g:go_metalinter_command = 'golangci-lint'
let g:go_metalinter_enabled = ['vet', 'golint', 'errcheck']
```

### Pre-commit Hooks

`.pre-commit-config.yaml`:
```yaml
repos:
  - repo: local
    hooks:
      - id: go-lint
        name: golangci-lint
        entry: golangci-lint run
        language: system
        files: \.go$
```

### GitHub Actions

Linting is integrated into CI:

```yaml
lint:
  runs-on: ubuntu-latest
  steps:
  - uses: actions/checkout@v4
  - name: Set up Go
    uses: actions/setup-go@v5
    with:
      go-version: 1.24.x
  - name: golangci-lint
    uses: golangci/golangci-lint-action@v6
    with:
      version: latest
```

## Disabling Linters

### Inline Disabling

```go
//nolint:errcheck // Justified reason
file, _ := os.Open("file.txt")

//nolint:gosec // G104: We're intentionally ignoring this error
defer file.Close()

// Disable multiple linters
//nolint:errcheck,gosec
problematicCode()
```

### File-level Disabling

```go
//nolint:gocritic // This file needs special handling
package main
```

### Disable Specific Rules

```go
//nolint:gosec[G401] // MD5 is acceptable for non-crypto use
hash := md5.Sum(data)
```

## Best Practices

### Code Quality

- **Fix, don't disable**: Address issues rather than disabling linters
- **Justify exceptions**: Always provide reasons for nolint directives
- **Regular updates**: Keep golangci-lint updated
- **Team consistency**: Use same configuration across team

### Performance

- **Use --fast**: For quick feedback during development
- **Cache results**: Enable caching for faster subsequent runs
- **Parallel execution**: Use -j flag for parallel processing

```bash
# Fast development checking
golangci-lint run --fast

# Enable caching
golangci-lint cache clean  # Clear cache if needed
golangci-lint run          # Uses cache automatically

# Parallel processing
golangci-lint run -j 4
```

### Configuration Management

- **Version control**: Always commit `.golangci.yml`
- **Team agreement**: Discuss and agree on enabled linters
- **Gradual adoption**: Enable new linters incrementally
- **Document exceptions**: Maintain list of approved exclusions

## Troubleshooting

### Common Issues

**Timeout errors**:
```bash
# Increase timeout
golangci-lint run --timeout=10m
```

**Memory issues**:
```bash
# Reduce concurrent workers
golangci-lint run -j 1
```

**Module download issues**:
```yaml
# In .golangci.yml
run:
  modules-download-mode: readonly
```

### Debug Mode

```bash
# Verbose output
golangci-lint run -v

# Debug information
golangci-lint run --debug

# Show enabled linters
golangci-lint linters
```

### Version Compatibility

```bash
# Check version
golangci-lint version

# Install specific version
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2
```