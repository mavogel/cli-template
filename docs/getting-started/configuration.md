# Configuration

Learn how to configure your CLI application for different environments and use cases.

## Configuration Files

### Cobra Configuration

The CLI uses Cobra for command-line interface management. Key configuration files:

- `cmd/root.go`: Root command configuration
- `cmd/*.go`: Individual command implementations

### Linting Configuration

`.golangci.yml` configures code quality checks:

```yaml
run:
  timeout: 5m
  modules-download-mode: readonly

linters:
  enable:
    - errcheck
    - gosimple
    - govet
    # ... more linters
```

### Release Configuration

`.goreleaser.yaml` configures automated releases:

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
```

## Environment Variables

### Development

```bash
# Enable debug mode
export DEBUG=true

# Set log level
export LOG_LEVEL=debug

# Custom config path
export CONFIG_PATH=/path/to/config
```

### Build Variables

GoReleaser injects build-time variables:

```go
var (
    version = "dev"
    commit  = "none"
    date    = "unknown"
)
```

## Command Configuration

### Adding Flags

```go
// Global flags (available to all commands)
func init() {
    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
    rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose output")
}

// Command-specific flags
func init() {
    helloCmd.Flags().StringP("name", "n", "", "Name to greet")
    helloCmd.Flags().IntP("count", "c", 1, "Number of greetings")
}
```

### Configuration Binding

Use Viper for advanced configuration management:

```go
import "github.com/spf13/viper"

func initConfig() {
    if cfgFile != "" {
        viper.SetConfigFile(cfgFile)
    } else {
        viper.AddConfigPath("$HOME")
        viper.SetConfigName(".cli-template")
    }
    
    viper.AutomaticEnv()
    
    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}
```

## Docker Configuration

### Dockerfile

```dockerfile
FROM scratch
COPY cli-template /usr/local/bin/cli-template
ENTRYPOINT ["/usr/local/bin/cli-template"]
```

### Docker Compose

```yaml
version: '3.8'
services:
  cli-template:
    build: .
    volumes:
      - ./config:/config
    environment:
      - CONFIG_PATH=/config/app.yaml
```

## GitHub Actions Configuration

### CI Configuration

`.github/workflows/ci.yml`:

```yaml
on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    strategy:
      matrix:
        go-version: [1.21.x, 1.22.x, 1.23.x]
```

### Release Configuration

`.github/workflows/release.yml`:

```yaml
on:
  push:
    tags:
      - 'v*'

permissions:
  contents: write
  packages: write
```

## Documentation Configuration

### MkDocs

`mkdocs.yml` configures documentation generation:

```yaml
site_name: CLI Template
theme:
  name: material
  features:
    - navigation.tabs
    - search.highlight
```

### Docker for Documentation

MkDocs is run via Docker container for consistency:

```bash
# Install Docker
brew install --cask docker

# Serve documentation locally
make docs-serve

# Build documentation
make docs-build
```

## Best Practices

### Security

- Never commit secrets to version control
- Use environment variables for sensitive data
- Implement proper input validation
- Follow the principle of least privilege

### Performance

- Use build constraints for platform-specific code
- Optimize binary size with build flags
- Implement proper error handling
- Use structured logging

### Maintainability

- Keep configuration files well-documented
- Use consistent naming conventions
- Implement feature flags for gradual rollouts
- Regular dependency updates