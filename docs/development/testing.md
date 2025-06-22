# Testing

Comprehensive testing strategies and best practices for your CLI application.

## Test Organization

### Test Structure

Tests are co-located with the code they test:

```
cmd/
├── root.go
├── root_test.go      # Tests for root command
├── hello.go  
└── hello_test.go     # Tests for hello command
```

### Test Naming

Follow Go testing conventions:

- **Files**: `*_test.go`
- **Functions**: `TestFunctionName`
- **Benchmarks**: `BenchmarkFunctionName`
- **Examples**: `ExampleFunctionName`

## Running Tests

### Basic Test Commands

```bash
# Run all tests
make test
go test ./...

# Run tests with verbose output
go test -v ./...

# Run specific package tests
go test ./cmd

# Run specific test function
go test -run TestHelloCommand ./cmd
```

### Test Coverage

```bash
# Generate coverage report
make test-coverage

# View coverage in terminal
go test -cover ./...

# Generate detailed coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# View coverage by function
go tool cover -func=coverage.out
```

## Writing Tests

### Command Testing Pattern

```go
func TestHelloCommand(t *testing.T) {
    tests := []struct {
        name     string
        args     []string
        expected string
    }{
        {
            name:     "default greeting",
            args:     []string{},
            expected: "Hello, World!",
        },
        {
            name:     "custom name",
            args:     []string{"--name", "Alice"},
            expected: "Hello, Alice!",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            buf := new(bytes.Buffer)
            
            cmd := &cobra.Command{
                Use: "hello",
                Run: func(cmd *cobra.Command, args []string) {
                    name, _ := cmd.Flags().GetString("name")
                    if name == "" {
                        name = "World"
                    }
                    cmd.Printf("Hello, %s!\n", name)
                },
            }
            cmd.Flags().StringP("name", "n", "", "Name to greet")
            
            cmd.SetOut(buf)
            cmd.SetArgs(tt.args)
            
            err := cmd.Execute()
            if err != nil {
                t.Errorf("Execute() error = %v", err)
            }

            output := buf.String()
            if output != tt.expected+"\n" {
                t.Errorf("Expected output %q, got %q", tt.expected+"\n", output)
            }
        })
    }
}
```

### Error Testing

```go
func TestCommandErrors(t *testing.T) {
    tests := []struct {
        name        string
        args        []string
        expectError bool
        errorMsg    string
    }{
        {
            name:        "invalid flag",
            args:        []string{"--invalid-flag"},
            expectError: true,
            errorMsg:    "unknown flag",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cmd := yourCommand() // Create command instance
            cmd.SetArgs(tt.args)
            
            err := cmd.Execute()
            
            if tt.expectError {
                if err == nil {
                    t.Error("Expected error but got none")
                }
                if !strings.Contains(err.Error(), tt.errorMsg) {
                    t.Errorf("Expected error containing %q, got %q", tt.errorMsg, err.Error())
                }
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
            }
        })
    }
}
```

## Test Types

### Unit Tests

Test individual functions and methods:

```go
func TestFormatGreeting(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"empty name", "", "Hello, World!"},
        {"with name", "Alice", "Hello, Alice!"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := formatGreeting(tt.input)
            if result != tt.expected {
                t.Errorf("formatGreeting(%q) = %q, want %q", tt.input, result, tt.expected)
            }
        })
    }
}
```

### Integration Tests

Test command execution end-to-end:

```go
func TestFullCommandExecution(t *testing.T) {
    // Create temporary directory for test artifacts
    tmpDir := t.TempDir()
    
    // Execute command with file output
    cmd := exec.Command("./cli-template", "hello", "--output", tmpDir+"/greeting.txt")
    err := cmd.Run()
    if err != nil {
        t.Fatalf("Command failed: %v", err)
    }
    
    // Verify output file
    content, err := os.ReadFile(tmpDir + "/greeting.txt")
    if err != nil {
        t.Fatalf("Failed to read output file: %v", err)
    }
    
    expected := "Hello, World!\n"
    if string(content) != expected {
        t.Errorf("Expected file content %q, got %q", expected, string(content))
    }
}
```

### Benchmark Tests

Measure performance:

```go
func BenchmarkHelloCommand(b *testing.B) {
    cmd := &cobra.Command{
        Use: "hello",
        Run: func(cmd *cobra.Command, args []string) {
            cmd.Print("Hello, World!")
        },
    }
    
    for i := 0; i < b.N; i++ {
        buf := new(bytes.Buffer)
        cmd.SetOut(buf)
        cmd.SetArgs([]string{})
        cmd.Execute()
    }
}
```

## Test Utilities

### Test Helpers

Create reusable test utilities:

```go
// testutil/helpers.go
package testutil

import (
    "bytes"
    "testing"
    "github.com/spf13/cobra"
)

func ExecuteCommand(t *testing.T, cmd *cobra.Command, args []string) (string, error) {
    t.Helper()
    
    buf := new(bytes.Buffer)
    cmd.SetOut(buf)
    cmd.SetErr(buf)
    cmd.SetArgs(args)
    
    err := cmd.Execute()
    return buf.String(), err
}

func AssertOutput(t *testing.T, got, want string) {
    t.Helper()
    
    if got != want {
        t.Errorf("Output mismatch:\nGot:  %q\nWant: %q", got, want)
    }
}
```

### Mock Dependencies

Use interfaces for testable code:

```go
// Interface for external dependencies
type FileWriter interface {
    WriteFile(filename string, data []byte, perm os.FileMode) error
}

// Production implementation
type OSFileWriter struct{}
func (w OSFileWriter) WriteFile(filename string, data []byte, perm os.FileMode) error {
    return os.WriteFile(filename, data, perm)
}

// Test mock
type MockFileWriter struct {
    WrittenFiles map[string][]byte
}
func (w *MockFileWriter) WriteFile(filename string, data []byte, perm os.FileMode) error {
    if w.WrittenFiles == nil {
        w.WrittenFiles = make(map[string][]byte)
    }
    w.WrittenFiles[filename] = data
    return nil
}
```

## Continuous Integration

### GitHub Actions

The CI workflow runs tests automatically:

```yaml
test:
  runs-on: ubuntu-latest
  strategy:
    matrix:
      go-version: [1.21.x, 1.22.x, 1.23.x]
  
  steps:
  - uses: actions/checkout@v4
  - name: Set up Go
    uses: actions/setup-go@v5
    with:
      go-version: ${{ matrix.go-version }}
  - name: Run tests
    run: go test -v -race -coverprofile=coverage.out ./...
```

### Test Configuration

Environment variables for tests:

```bash
# Enable race detection
export GORACE="halt_on_error=1"

# Test timeout
export GOTESTSUM_TIMEOUT=300s

# Coverage threshold
export COVERAGE_THRESHOLD=80
```

## Best Practices

### Test Organization

- **Table-driven tests**: Use for multiple similar scenarios
- **Subtests**: Group related test cases with `t.Run()`
- **Helper functions**: Extract common test setup
- **Test fixtures**: Use consistent test data

### Test Reliability

- **Deterministic**: Tests should produce same results
- **Independent**: Tests shouldn't depend on each other
- **Fast**: Keep tests quick to encourage frequent running
- **Clear failures**: Provide helpful error messages

### Coverage Goals

- **Aim for 80%+**: Good coverage without obsessing over 100%
- **Test edge cases**: Error conditions, empty inputs, boundaries
- **Focus on critical paths**: Prioritize important functionality
- **Ignore generated code**: Exclude auto-generated files

### Common Patterns

```go
// Setup and teardown
func TestWithSetup(t *testing.T) {
    // Setup
    oldEnv := os.Getenv("TEST_VAR")
    os.Setenv("TEST_VAR", "test-value")
    defer os.Setenv("TEST_VAR", oldEnv) // Cleanup
    
    // Test implementation
}

// Parallel tests
func TestParallel(t *testing.T) {
    tests := []struct{...}{}
    
    for _, tt := range tests {
        tt := tt // Capture range variable
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel() // Run in parallel
            // Test implementation
        })
    }
}
```