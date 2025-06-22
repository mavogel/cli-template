# Contributing Guidelines

Welcome to the CLI Template project! We appreciate your interest in contributing.

## Getting Started

### Prerequisites

Before contributing, ensure you have:

- **Go 1.21+**: [Install Go](https://golang.org/dl/)
- **Git**: [Install Git](https://git-scm.com/downloads)
- **Make**: Usually available on Unix-like systems
- **golangci-lint**: For code quality checks
- **GoReleaser**: For release testing

### Development Setup

1. **Fork the repository** on GitHub
2. **Clone your fork**:
   ```bash
   git clone https://github.com/yourusername/cli-template.git
   cd cli-template
   ```
3. **Install dependencies**:
   ```bash
   go mod download
   make deps
   ```
4. **Verify setup**:
   ```bash
   make test
   make lint
   make build
   ```

## Development Workflow

### Branching Strategy

- **main**: Stable, production-ready code
- **develop**: Integration branch for features
- **feature/***: Feature development branches
- **bugfix/***: Bug fix branches
- **hotfix/***: Critical fixes for production

### Making Changes

1. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** following our coding standards

3. **Test your changes**:
   ```bash
   make test
   make lint
   make build
   ```

4. **Commit with conventional commits**:
   ```bash
   git commit -m "feat: add new command for data processing"
   ```

5. **Push and create a Pull Request**

## Code Standards

### Go Code Style

Follow standard Go conventions:

- **gofmt**: Use standard formatting
- **golint**: Follow linting recommendations  
- **go vet**: Pass static analysis
- **Comments**: Document public APIs

```go
// Good: Proper function documentation
// ProcessData processes the input data and returns the result.
// It returns an error if the data format is invalid.
func ProcessData(input []byte) ([]byte, error) {
    // Implementation
}
```

### Testing

- **Test coverage**: Aim for 80%+ coverage
- **Table-driven tests**: Use for multiple scenarios
- **Test naming**: `TestFunctionName_Scenario`
- **Mock external dependencies**: Use interfaces

```go
func TestProcessData_ValidInput(t *testing.T) {
    tests := []struct {
        name     string
        input    []byte
        expected []byte
        wantErr  bool
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

### Documentation

- **Code comments**: Document public functions
- **README updates**: For significant changes
- **Changelog**: Follow conventional changelog format
- **API docs**: Update for new commands/flags

## Commit Guidelines

### Conventional Commits

Use [Conventional Commits](https://conventionalcommits.org/) format:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

#### Types

- **feat**: New feature
- **fix**: Bug fix
- **docs**: Documentation changes
- **style**: Code style changes (formatting, etc.)
- **refactor**: Code refactoring
- **test**: Adding or updating tests
- **chore**: Maintenance tasks
- **ci**: CI/CD changes

#### Examples

```bash
# Feature
git commit -m "feat: add JSON output format support"

# Bug fix
git commit -m "fix: resolve panic when input file is empty"

# Documentation
git commit -m "docs: update installation instructions"

# Breaking change
git commit -m "feat!: change configuration file format to YAML"
```

## Pull Request Process

### Before Submitting

1. **Rebase on latest main**:
   ```bash
   git fetch origin
   git rebase origin/main
   ```

2. **Run full test suite**:
   ```bash
   make all
   ```

3. **Update documentation** if needed

4. **Add/update tests** for new functionality

### PR Template

Use this template for Pull Requests:

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
- [ ] Tests pass locally
- [ ] Added tests for new functionality
- [ ] Manual testing performed

## Checklist
- [ ] Code follows style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] No new warnings introduced
```

### Review Process

1. **Automated checks** must pass
2. **Code review** by maintainers
3. **Manual testing** if required
4. **Approval** and merge

## Issue Guidelines

### Bug Reports

Use the bug report template:

```markdown
**Describe the bug**
Clear description of the bug

**To Reproduce**
Steps to reproduce:
1. Run command '...'
2. See error

**Expected behavior**
What should happen

**Environment**
- OS: [e.g., Ubuntu 20.04]
- Go version: [e.g., 1.21]
- CLI version: [e.g., v1.0.0]
```

### Feature Requests

Use the feature request template:

```markdown
**Feature Description**
Clear description of the feature

**Use Case**
Why is this feature needed?

**Proposed Solution**
How should it work?

**Alternatives**
Other solutions considered
```

## Community Guidelines

### Code of Conduct

- **Be respectful**: Treat everyone with respect
- **Be inclusive**: Welcome newcomers
- **Be constructive**: Provide helpful feedback
- **Be patient**: Help others learn

### Communication

- **Issues**: For bugs and feature requests
- **Discussions**: For questions and ideas
- **Pull Requests**: For code contributions
- **Email**: For security issues

## Release Process

### Version Numbering

Follow [Semantic Versioning](https://semver.org/):
- **MAJOR**: Breaking changes
- **MINOR**: New features (backward compatible)
- **PATCH**: Bug fixes

### Release Workflow

1. **Create release branch**: `release/v1.0.0`
2. **Update version**: In relevant files
3. **Update changelog**: Document changes
4. **Test thoroughly**: Manual and automated
5. **Merge to main**: Via pull request
6. **Tag release**: `git tag v1.0.0`
7. **Push tag**: Triggers automated release

## Recognition

### Contributors

All contributors are recognized in:
- **CONTRIBUTORS.md**: List of contributors
- **Release notes**: Major contributions highlighted
- **GitHub**: Contributor graph and stats

### Maintainers

Current maintainers:
- @mavogel - Project lead

## Getting Help

### Resources

- **Documentation**: [docs.example.com](https://mavogel.github.io/cli-template/)
- **Issues**: GitHub Issues
- **Discussions**: GitHub Discussions

### Support Channels

- **Bug reports**: GitHub Issues
- **Feature requests**: GitHub Issues  
- **Questions**: GitHub Discussions
- **Security**: Email maintainers

## License

By contributing, you agree that your contributions will be licensed under the same license as the project.