# Project Structure

Understanding the CLI Template project layout and organization.

## Directory Overview

```
cli-template/
├── .github/                # GitHub specific files
│   └── workflows/          # GitHub Actions workflows
│       ├── ci.yml         # Continuous integration
│       └── release.yml    # Release automation
├── cmd/                   # Command implementations
│   ├── root.go           # Root command definition
│   ├── hello.go          # Example subcommand
│   ├── root_test.go      # Root command tests
│   └── hello_test.go     # Hello command tests
├── docs/                 # Documentation source files
│   ├── index.md         # Homepage
│   ├── getting-started/ # Getting started guides
│   ├── development/     # Development documentation
│   ├── deployment/      # Deployment guides
│   ├── reference/       # API and command reference
│   └── contributing/    # Contribution guidelines
├── dist/                # Build artifacts (generated)
├── bin/                 # Local build output
├── .golangci.yml       # Linting configuration
├── .goreleaser.yaml    # Release configuration
├── Dockerfile          # Container build instructions
├── Makefile           # Build automation
├── mkdocs.yml         # Documentation configuration
├── go.mod             # Go module definition
├── go.sum             # Go module checksums
└── main.go            # Application entry point
```

## Core Components

### Application Entry Point

**`main.go`**
- Application bootstrap
- Version information injection
- Error handling and exit codes

```go
func main() {
    cmd.SetVersionInfo(version, commit, date)
    if err := cmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
```

### Command Structure

**`cmd/` Directory**
- Each command is a separate file
- Tests are co-located with commands
- Follows Cobra command patterns

**`cmd/root.go`**
- Defines the root command
- Global flags and configuration
- Version information display

**`cmd/hello.go`**
- Example subcommand implementation
- Demonstrates flag usage
- Shows command structure patterns
- Separates business logic into testable action functions

### Build Configuration

**`Makefile`**
- Standardized build tasks
- Development workflow automation
- Cross-platform compatibility

**`.goreleaser.yaml`**
- Multi-platform build configuration
- Release artifact generation
- Distribution automation

### Quality Assurance

**`.golangci.yml`**
- Comprehensive linting rules
- Code quality enforcement
- Security checks

**`*_test.go` Files**
- Unit tests for each component
- Table-driven test patterns
- Coverage reporting

## File Organization Principles

### Commands (`cmd/`)

Each command follows this pattern:

```go
// cmd/example.go
package cmd

import (
    "github.com/spf13/cobra"
)

var exampleCmd = &cobra.Command{
    Use:   "example",
    Short: "Example command description",
    Long:  `Detailed example command description.`,
    Run:   runExample,
}

func runExample(cmd *cobra.Command, args []string) {
    // Command implementation
}

func init() {
    rootCmd.AddCommand(exampleCmd)
    // Add flags here
}
```

### Tests (`*_test.go`)

Test files follow Go conventions:

```go
// cmd/example_test.go
package cmd

import (
    "testing"
)

func TestExampleCommand(t *testing.T) {
    tests := []struct {
        name     string
        args     []string
        expected string
    }{
        // Test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

## Configuration Files

### Linting (`.golangci.yml`)

Organized by categories:
- **Run settings**: Timeout, module handling
- **Enabled linters**: Code quality tools
- **Linter settings**: Individual tool configuration
- **Issue filtering**: Test-specific exclusions

### Release (`.goreleaser.yaml`)

Structured for automation:
- **Before hooks**: Pre-build preparation
- **Builds**: Multi-platform compilation
- **Archives**: Distribution packaging
- **Release**: GitHub integration

### Documentation (`mkdocs.yml`)

Content organization:
- **Site metadata**: Name, description, URLs
- **Theme configuration**: Material design
- **Navigation structure**: Hierarchical content
- **Plugin configuration**: Additional features

## Development Workflow

### Adding New Commands

1. Create `cmd/newcommand.go`
2. Implement command structure
3. Add to root command in `init()`
4. Create `cmd/newcommand_test.go`
5. Update documentation

### Modifying Build Process

1. Update `Makefile` for new tasks
2. Modify `.goreleaser.yaml` for releases
3. Adjust CI workflows if needed
4. Test with `make all`

### Documentation Updates

1. Edit relevant `.md` files in `docs/`
2. Update `mkdocs.yml` navigation
3. Test locally with `mkdocs serve`
4. Commit changes for auto-deployment

## Best Practices

### Code Organization

- **Single responsibility**: One command per file
- **Clear naming**: Descriptive file and function names
- **Consistent structure**: Follow established patterns
- **Proper imports**: Group standard, third-party, and local

### Testing Strategy

- **Unit tests**: Test individual functions
- **Integration tests**: Test command execution
- **Table-driven tests**: Multiple scenarios
- **Error cases**: Test failure conditions

### Documentation

- **Code comments**: Document public APIs
- **README files**: Overview and quick start
- **Detailed docs**: Comprehensive guides
- **Examples**: Real-world usage patterns