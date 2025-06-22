# GitHub Actions

Continuous Integration and Deployment workflows for automated testing, building, and releasing.

## Workflow Overview

The project includes two main GitHub Actions workflows:

1. **CI Workflow** (`.github/workflows/ci.yml`) - Continuous Integration
2. **Release Workflow** (`.github/workflows/release.yml`) - Automated Releases

## CI Workflow

### Triggers

The CI workflow runs on:
- **Push** to `main` and `develop` branches
- **Pull requests** targeting `main` branch

```yaml
on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]
```

### Jobs

#### Test Job

Runs tests across multiple Go versions:

```yaml
test:
  runs-on: ubuntu-latest
  strategy:
    matrix:
      go-version: [1.21.x, 1.22.x, 1.23.x]
```

**Steps**:
1. Checkout code
2. Setup Go environment
3. Cache Go modules
4. Download dependencies
5. Run tests with coverage
6. Upload coverage to Codecov

#### Lint Job

Code quality checks:

```yaml
lint:
  runs-on: ubuntu-latest
  steps:
  - name: golangci-lint
    uses: golangci/golangci-lint-action@v6
```

#### Build Job

Verify the application builds:

```yaml
build:
  runs-on: ubuntu-latest
  needs: [test, lint]
  steps:
  - name: Build
    run: go build -v ./...
```

#### GoReleaser Check

Validates release configuration:

```yaml
goreleaser-check:
  runs-on: ubuntu-latest
  needs: [test, lint]
  steps:
  - name: Check GoReleaser config
    uses: goreleaser/goreleaser-action@v6
    with:
      args: check
```

## Release Workflow

### Triggers

The release workflow triggers on version tags:

```yaml
on:
  push:
    tags:
      - 'v*'
```

### Permissions

Required permissions for release automation:

```yaml
permissions:
  contents: write    # Create releases
  packages: write    # Push container images
```

### Release Job

Single job that handles the complete release process:

1. **Checkout** with full history (`fetch-depth: 0`)
2. **Setup Go** environment
3. **Login** to GitHub Container Registry
4. **Run GoReleaser** with release configuration

```yaml
- name: Run GoReleaser
  uses: goreleaser/goreleaser-action@v6
  with:
    distribution: goreleaser
    version: latest
    args: release --clean
  env:
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

## Workflow Features

### Caching

Go module caching for faster builds:

```yaml
- name: Cache Go modules
  uses: actions/cache@v4
  with:
    path: ~/go/pkg/mod
    key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
    restore-keys: |
      ${{ runner.os }}-go-
```

### Matrix Testing

Tests across multiple Go versions:

```yaml
strategy:
  matrix:
    go-version: [1.21.x, 1.22.x, 1.23.x]
```

### Job Dependencies

Ensures proper execution order:

```yaml
build:
  needs: [test, lint]  # Only run if test and lint pass
```

## Secrets and Variables

### Required Secrets

- **GITHUB_TOKEN**: Automatically provided by GitHub
- **CODECOV_TOKEN**: Optional for coverage reporting

### Environment Variables

Available in workflows:

```yaml
env:
  GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
  GO_VERSION: 1.23.x
```

## Customization

### Adding New Jobs

Add custom jobs to CI workflow:

```yaml
security-scan:
  runs-on: ubuntu-latest
  steps:
  - uses: actions/checkout@v4
  - name: Run security scan
    uses: securecodewarrior/github-action-add-sarif@v1
    with:
      sarif-file: security-results.sarif
```

### Custom Build Steps

Extend the build process:

```yaml
- name: Custom build step
  run: |
    echo "Running custom build logic"
    make custom-target
```

### Conditional Execution

Run steps based on conditions:

```yaml
- name: Deploy to staging
  if: github.ref == 'refs/heads/develop'
  run: make deploy-staging
```

## Monitoring

### Workflow Status

Check workflow status:

```bash
# List recent runs
gh run list

# View specific run
gh run view <run-id>

# Watch live run
gh run watch
```

### Notifications

Configure notifications for failed workflows:

1. **Repository Settings** → **Notifications**
2. **Actions** → **Failed workflows**
3. **Email/Slack integration**

## Best Practices

### Security

- **Minimal permissions**: Only grant necessary permissions
- **Secret management**: Use GitHub secrets for sensitive data
- **Third-party actions**: Pin to specific versions
- **Code scanning**: Enable security scanning

```yaml
- uses: actions/checkout@v4  # Pinned version
  with:
    token: ${{ secrets.GITHUB_TOKEN }}
```

### Performance

- **Caching**: Cache dependencies and build artifacts
- **Parallel jobs**: Run independent jobs in parallel
- **Conditional runs**: Skip unnecessary steps
- **Resource limits**: Use appropriate runner sizes

### Reliability

- **Timeouts**: Set appropriate timeouts
- **Retry logic**: Retry flaky operations
- **Fallback strategies**: Handle external service failures
- **Status checks**: Require passing checks for merges

```yaml
- name: Test with retry
  uses: nick-invision/retry@v2
  with:
    timeout_minutes: 10
    max_attempts: 3
    command: make test
```

## Troubleshooting

### Common Issues

**Permission denied errors**:
```yaml
permissions:
  contents: write
  packages: write
```

**Cache misses**:
```yaml
- name: Debug cache
  run: |
    echo "Cache key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}"
    ls -la ~/go/pkg/mod || echo "No cache found"
```

**GoReleaser failures**:
```bash
# Local testing
make release-check
make release-snapshot
```

### Debugging Workflows

Enable debug logging:

1. **Repository Settings** → **Actions** → **General**
2. **Enable debug logging**
3. **Re-run workflow**

Or add debug steps:

```yaml
- name: Debug environment
  run: |
    echo "Runner OS: ${{ runner.os }}"
    echo "GitHub ref: ${{ github.ref }}"
    echo "Event name: ${{ github.event_name }}"
    env | sort
```

### Log Analysis

View detailed logs:

```bash
# Download logs
gh run download <run-id>

# View job logs
gh run view <run-id> --log
```

## Advanced Configuration

### Custom Runners

Use self-hosted runners:

```yaml
runs-on: [self-hosted, linux, x64]
```

### Workflow Templates

Create reusable workflows:

```yaml
# .github/workflows/reusable-test.yml
on:
  workflow_call:
    inputs:
      go-version:
        required: true
        type: string

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/setup-go@v5
      with:
        go-version: ${{ inputs.go-version }}
```

### Marketplace Actions

Leverage community actions:

```yaml
- name: Setup Go
  uses: actions/setup-go@v5
  
- name: golangci-lint
  uses: golangci/golangci-lint-action@v6
  
- name: GoReleaser
  uses: goreleaser/goreleaser-action@v6
```