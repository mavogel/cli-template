# Releases

Automated release management with GoReleaser and GitHub Actions.

## Release Process

### Automated Releases

The project uses GoReleaser for automated, cross-platform releases triggered by Git tags:

1. **Tag a version**: `git tag v1.0.0`
2. **Push the tag**: `git push origin v1.0.0`
3. **GitHub Actions**: Automatically builds and releases
4. **Artifacts**: Binaries, archives, and checksums are published

### Manual Releases

For testing or custom releases:

```bash
# Validate configuration
make release-check

# Create snapshot (testing)
make release-snapshot

# Manual release (requires GITHUB_TOKEN)
export GITHUB_TOKEN=your_token
make release
```

## Version Strategy

### Semantic Versioning

Follow [semantic versioning](https://semver.org/) principles:

- **MAJOR**: Breaking changes (`v1.0.0` → `v2.0.0`)
- **MINOR**: New features, backward compatible (`v1.0.0` → `v1.1.0`)
- **PATCH**: Bug fixes (`v1.0.0` → `v1.0.1`)

### Pre-releases

Use pre-release identifiers for testing:

```bash
# Alpha release
git tag v1.0.0-alpha.1

# Beta release  
git tag v1.0.0-beta.1

# Release candidate
git tag v1.0.0-rc.1
```

## Release Artifacts

### Binary Distributions

GoReleaser creates optimized binaries for multiple platforms:

| Platform | Architecture | Binary Name |
|----------|-------------|-------------|
| Linux | amd64 | `cli-template_linux_amd64` |
| Linux | arm64 | `cli-template_linux_arm64` |
| macOS | amd64 | `cli-template_darwin_amd64` |
| macOS | arm64 | `cli-template_darwin_arm64` |
| Windows | amd64 | `cli-template_windows_amd64.exe` |

### Archives

Platform-specific archives with consistent naming:

- **Linux/macOS**: `.tar.gz` format
- **Windows**: `.zip` format
- **Naming**: `cli-template_OS_ARCHITECTURE.ext`

### Checksums

SHA256 checksums for all artifacts:
- **File**: `checksums.txt`
- **Usage**: Verify download integrity

```bash
# Verify download
sha256sum -c checksums.txt
```

## Build Configuration

### GoReleaser Settings

Key configuration in `.goreleaser.yaml`:

```yaml
builds:
  - main: ./main.go
    binary: cli-template
    env:
      - CGO_ENABLED=0  # Static builds
    goos: [linux, windows, darwin]
    goarch: [amd64, arm64]
    ldflags:
      - -s -w  # Strip debug info
      - -X main.version={{.Version}}
      - -X main.commit={{.Commit}}
      - -X main.date={{.Date}}
```

### Version Injection

Build-time variables are injected automatically:

```go
var (
    version = "dev"     // Replaced with git tag
    commit  = "none"    // Replaced with git commit
    date    = "unknown" // Replaced with build date
)
```

View version information:
```bash
cli-template --version
```

## Distribution Channels

### GitHub Releases

Primary distribution through GitHub Releases:
- **Automatic**: Triggered by tags
- **Draft mode**: Review before publishing
- **Release notes**: Generated from commits

### Homebrew

Automated Homebrew tap updates:

```yaml
brews:
  - name: cli-template
    description: "A CLI application template"
    homepage: "https://github.com/mavogel/cli-template"
    repository:
      owner: mavogel
      name: homebrew-tap
```

Install via Homebrew:
```bash
brew tap mavogel/homebrew-tap
brew install cli-template
```

### Container Registry

Docker images published to GitHub Container Registry:

```bash
# Pull latest
docker pull ghcr.io/mavogel/cli-template:latest

# Pull specific version
docker pull ghcr.io/mavogel/cli-template:v1.0.0

# Run container
docker run --rm ghcr.io/mavogel/cli-template:latest --help
```

## Release Notes

### Changelog Generation

Automatic changelog from commit messages:

```yaml
changelog:
  sort: asc
  use: github
  groups:
    - title: Features
      regexp: '^.*?feat(\([[:word:]]+\))??!?:.+$'
    - title: 'Bug fixes'
      regexp: '^.*?fix(\([[:word:]]+\))??!?:.+$'
```

### Commit Message Format

Use conventional commits for better changelog:

```bash
# Features
git commit -m "feat: add new command for data processing"

# Bug fixes
git commit -m "fix: resolve parsing error in config loader"

# Breaking changes
git commit -m "feat!: change API for configuration handling"
```

## Release Workflow

### Development Cycle

1. **Feature development** on feature branches
2. **Pull request** review and merge to main
3. **Version tag** when ready for release
4. **Automated release** via GitHub Actions
5. **Distribution** to all channels

### Release Checklist

Before tagging a release:

- [ ] All tests passing
- [ ] Documentation updated
- [ ] CHANGELOG.md reviewed
- [ ] Breaking changes documented
- [ ] Version number follows semantic versioning

```bash
# Pre-release validation
make all
make release-check
```

### Tag Creation

```bash
# Create and push tag
git tag -a v1.0.0 -m "Release version 1.0.0"
git push origin v1.0.0

# Or use GitHub CLI
gh release create v1.0.0 --generate-notes
```

## Troubleshooting

### Failed Releases

Common issues and solutions:

**GoReleaser validation fails**:
```bash
# Check configuration
make release-check

# Validate specific sections
goreleaser check --debug
```

**GitHub token issues**:
```bash
# Check token permissions
gh auth status

# Re-authenticate
gh auth login
```

**Build failures**:
```bash
# Test local build
make release-snapshot

# Check build logs in GitHub Actions
gh run list
gh run view <run-id>
```

### Version Conflicts

**Tag already exists**:
```bash
# Delete local tag
git tag -d v1.0.0

# Delete remote tag
git push origin --delete v1.0.0

# Create new tag
git tag v1.0.1
```

**Pre-release cleanup**:
```bash
# Delete draft release
gh release delete v1.0.0

# Re-run release
git push origin v1.0.0
```

## Best Practices

### Release Timing

- **Regular cadence**: Monthly or bi-weekly releases
- **Feature completion**: Don't rush incomplete features
- **Security fixes**: Immediate patch releases
- **Breaking changes**: Major version bumps

### Communication

- **Release notes**: Clear, user-focused descriptions
- **Migration guides**: For breaking changes
- **Deprecation notices**: Advance warning for removals
- **Security advisories**: For vulnerability fixes

### Quality Assurance

- **Testing**: Comprehensive test suite before release
- **Staging**: Test snapshot builds
- **Rollback plan**: Ability to revert if needed
- **Monitoring**: Watch for issues post-release