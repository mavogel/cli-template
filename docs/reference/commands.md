# Commands Reference

Complete reference for all available CLI commands and their options.

## Root Command

### cli-template

The main entry point for the CLI application.

```bash
cli-template [flags]
cli-template [command]
```

**Global Flags:**
- `-h, --help`: Show help information
- `-v, --version`: Display version information

**Available Commands:**
- `hello`: Print a greeting message
- `completion`: Generate shell completion scripts
- `help`: Show help for any command

## Commands

### hello

Print a greeting message with optional customization.

```bash
cli-template hello [flags]
```

**Usage Examples:**
```bash
# Basic greeting
cli-template hello

# Custom name
cli-template hello --name "Alice"
cli-template hello -n "Bob"
```

**Flags:**
- `-n, --name string`: Name to greet (default: "World")
- `-h, --help`: Help for hello command

### completion

Generate shell completion scripts for various shells.

```bash
cli-template completion [bash|zsh|fish|powershell]
```

**Usage Examples:**
```bash
# Bash completion
cli-template completion bash > /etc/bash_completion.d/cli-template

# Zsh completion
cli-template completion zsh > "${fpath[1]}/_cli-template"

# Fish completion
cli-template completion fish > ~/.config/fish/completions/cli-template.fish

# PowerShell completion
cli-template completion powershell > cli-template.ps1
```

## Exit Codes

The CLI application uses standard exit codes:

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Invalid usage/arguments |
| 130 | Interrupted by user (Ctrl+C) |

## Environment Variables

### Configuration

- `CONFIG_PATH`: Path to configuration file
- `LOG_LEVEL`: Logging level (debug, info, warn, error)
- `DEBUG`: Enable debug mode (true/false)

### Build Information

These are set at build time:
- `VERSION`: Application version
- `COMMIT`: Git commit hash
- `BUILD_DATE`: Build timestamp

## Global Configuration

### Config File Locations

The CLI looks for configuration files in:

1. Current directory: `./cli-template.yaml`
2. Home directory: `~/.cli-template.yaml`
3. System config: `/etc/cli-template/config.yaml`

### Configuration Format

```yaml
# Example configuration
log_level: info
output_format: text
timeout: 30s

# Command-specific settings
hello:
  default_name: "World"
  greeting_format: "Hello, %s!"
```

## Advanced Usage

### Shell Integration

#### Bash

Add to `~/.bashrc`:
```bash
# Enable completion
source <(cli-template completion bash)

# Add alias
alias ct='cli-template'
```

#### Zsh

Add to `~/.zshrc`:
```bash
# Enable completion
autoload -U compinit; compinit
source <(cli-template completion zsh)

# Add alias
alias ct='cli-template'
```

#### Fish

Add to `~/.config/fish/config.fish`:
```fish
# Enable completion
cli-template completion fish | source

# Add alias
alias ct='cli-template'
```

### Scripting

Use in shell scripts:

```bash
#!/bin/bash
set -e

# Check if cli-template is available
if ! command -v cli-template &> /dev/null; then
    echo "cli-template not found"
    exit 1
fi

# Use in script
result=$(cli-template hello --name "Script")
echo "Result: $result"

# Check exit code
if cli-template hello --name "Test"; then
    echo "Command succeeded"
else
    echo "Command failed with exit code $?"
fi
```

### JSON Output

Enable JSON output for programmatic use:

```bash
# Example with structured output
cli-template hello --name "Alice" --format json
```

Output:
```json
{
  "message": "Hello, Alice!",
  "timestamp": "2023-12-07T10:30:00Z",
  "status": "success"
}
```