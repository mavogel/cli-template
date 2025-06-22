# Quick Start

Get up and running with CLI Template in minutes.

## Basic Usage

### Hello Command

The template includes a sample `hello` command:

```bash
# Basic greeting
cli-template hello
# Output: Hello, World!

# Custom greeting
cli-template hello --name "Developer"
# Output: Hello, Developer!

# Short flag
cli-template hello -n "Go"
# Output: Hello, Go!
```

### Help and Version

```bash
# Show help
cli-template --help
cli-template hello --help

# Show version
cli-template --version
```

## Customizing Your CLI

### 1. Modify the Root Command

Edit `cmd/root.go` to customize the main command:

```go
var rootCmd = &cobra.Command{
    Use:   "your-app-name",
    Short: "Your app description",
    Long:  `Your detailed app description.`,
    // ... rest of configuration
}
```

### 2. Add New Commands

Create new command files in the `cmd/` directory:

```go
// cmd/your-command.go
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var yourCmd = &cobra.Command{
    Use:   "your-command",
    Short: "Description of your command",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Your command executed!")
    },
}

func init() {
    rootCmd.AddCommand(yourCmd)
}
```

### 3. Update Module Name

Update the Go module name in `go.mod`:

```go
module github.com/yourusername/your-app-name

go 1.21
```

And update all imports accordingly.

## Building Your Application

### Development Build

```bash
# Build for current platform
make build

# Run without building
make run

# Development mode with hot reload
make dev
```

### Cross-Platform Build

```bash
# Create snapshot release (all platforms)
make release-snapshot

# Check built binaries
ls dist/
```

### Production Release

```bash
# Tag a version
git tag v1.0.0
git push origin v1.0.0

# This triggers GitHub Actions to create a release
```

## Testing

### Run Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# View coverage report
open coverage.html
```

### Add New Tests

Create test files alongside your commands:

```go
// cmd/your-command_test.go
package cmd

import (
    "testing"
)

func TestYourCommand(t *testing.T) {
    // Your test implementation
}
```

## Next Steps

- [Project Structure](../development/project-structure.md) - Understand the codebase
- [Configuration](configuration.md) - Configure your application
- [Development Guide](../development/building.md) - Advanced development topics