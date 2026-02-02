# Pre-commit Hooks

Podgrab uses [pre-commit](https://pre-commit.com/) to automatically check code
quality before commits.

## Quick Setup

```bash
# Install pre-commit
brew install pre-commit  # macOS
# or: pip install pre-commit

# Install hooks (from project root)
pre-commit install
pre-commit install --hook-type commit-msg

# Verify setup
pre-commit run --all-files
```

## What Gets Checked

### Go Code (Primary Focus)

- **Formatting**: `gofmt -s`, `goimports` (auto-fix)
- **Linting**: 20+ linters via golangci-lint
- **Build**: Code must compile (`go build`)
- **Vet**: `go vet` checks
- **Complexity**: Functions \<15 cyclomatic complexity
- **Security**: gosec vulnerability scanner
- **Dependencies**: `go mod tidy` (auto-fix)

### Files & Formats

- **Markdown**: mdformat with GFM support (auto-fix, wrap=80)
- **YAML**: Syntax validation and linting
- **JSON/TOML**: Syntax checking
- **Dockerfile**: hadolint best practices
- **Shell scripts**: shellcheck warnings

### General Checks

- File size (\<10MB)
- Secrets detection (API keys, passwords)
- Trailing whitespace (auto-fix)
- Line endings (LF, auto-fix)
- Merge conflicts
- Private keys

## Usage

### Automatic (Default)

Pre-commit runs automatically on `git commit`:

```bash
git add main.go
git commit -m "feat: add feature"
# Hooks run automatically
```

### Manual Execution

```bash
# Run all checks
pre-commit run --all-files

# Run specific hook
pre-commit run golangci-lint --all-files

# Run on specific files
pre-commit run --files main.go service/podcastService.go
```

### Skip Hooks (Emergency Only)

```bash
# Skip all hooks (not recommended)
git commit --no-verify -m "Emergency fix"

# Skip specific slow hooks
SKIP=golangci-lint,gosec git commit -m "WIP"
```

## Configuration

### Main Config: `.pre-commit-config.yaml`

Defines all hooks. Key sections:

```yaml
repos:
  # Go tools (primary)
  - repo: github.com/dnephin/pre-commit-golang
    hooks: [go-fmt, go-imports, go-vet, etc.]

  # Comprehensive linting
  - repo: github.com/golangci/golangci-lint
    hooks: [golangci-lint]

  # Security
  - repo: github.com/securego/gosec
    hooks: [gosec]
```

### Go Linting: `.golangci.yml`

Controls 20+ Go linters:

```yaml
linters:
  enable:
    - errcheck, gosimple, govet
    - staticcheck, unused
    - gosec, bodyclose, dupl
    # ... 15+ more

linters-settings:
  gocyclo:
    min-complexity: 15
```

**To adjust complexity threshold:**

```yaml
gocyclo:
  min-complexity: 20  # Default: 15
```

### YAML Rules: `.yamllint.yml`

```yaml
rules:
  line-length:
    max: 120
  indentation:
    spaces: 2
```

## Updating Hooks

Pre-commit hooks are versioned and should be updated regularly:

```bash
# Update all hooks (uses --freeze for reproducibility)
pre-commit autoupdate --freeze

# Or use the project update script
./scripts/update-project.sh
```

## Common Issues

### "golangci-lint not found"

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### "Pre-commit too slow"

```bash
# Skip expensive checks during development
SKIP=golangci-lint,gosec git commit -m "WIP"

# Run full checks before pushing
pre-commit run --all-files
```

### "Commit message rejected"

Use conventional commit format:

```bash
# Good
git commit -m "feat: add dark mode"
git commit -m "fix: resolve download issue"
git commit -m "docs: update API docs"

# Bad - missing type prefix
git commit -m "added feature"
```

**Types**: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `chore`,
`ci`

### False Positives

Disable specific linter for a line:

```go
// nolint:gosec // G304: File path from trusted source
file, err := os.Open(userInput)
```

Or exclude in `.golangci.yml`:

```yaml
issues:
  exclude-rules:
    - linters: [gosec]
      text: "G304"
```

## Hook Performance

| Hook          | Speed     | Auto-fix   | Notes           |
| ------------- | --------- | ---------- | --------------- |
| go-fmt        | ⚡ Fast   | ✅ Yes     | Always runs     |
| go-imports    | ⚡ Fast   | ✅ Yes     | Always runs     |
| go-vet        | 🐢 Medium | ❌ No      | Always runs     |
| golangci-lint | 🐌 Slow   | ✅ Partial | Can skip in dev |
| gosec         | 🐢 Medium | ❌ No      | Can skip in dev |
| mdformat      | ⚡ Fast   | ✅ Yes     | Always runs     |

**First run**: Slow (installs hook environments, ~2-5 minutes) **Subsequent
runs**: Fast (cached, ~10-30 seconds)

## CI Integration

To run pre-commit in GitHub Actions:

```yaml
# .github/workflows/pre-commit.yml
name: pre-commit
on: [push, pull_request]
jobs:
  pre-commit:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-python@v5
      - uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      - uses: pre-commit/action@v3.0.0
```

## Best Practices

**Do** ✅:

- Run `pre-commit run --all-files` before pushing
- Fix issues shown by hooks
- Update hooks regularly with `pre-commit autoupdate --freeze`
- Use conventional commit messages

**Don't** ❌:

- Use `--no-verify` unless absolutely necessary
- Ignore security warnings without understanding them
- Skip checks just to save time
- Commit work-in-progress to main branch

## Project-Specific Notes

### Go Focus

This pre-commit configuration prioritizes Go tooling:

- Go hooks run first and are most comprehensive
- Other language checks (YAML, Markdown) are secondary
- Security scanning focuses on Go vulnerabilities

### Markdown Formatting

- Uses `mdformat` with GFM (GitHub Flavored Markdown)
- Auto-wraps at 80 characters
- Formats tables consistently
- Excludes `CHANGELOG.md` (manually managed)

### Excluded Paths

Pre-commit skips these directories:

- `webassets/` - Static web assets
- `client/` - HTML templates
- `vendor/` - Go vendor directory
- `assets/` - Downloaded podcast files

## Related Documentation

- [Contributing Guide](contributing.md) - Full contribution workflow
- [Development Setup](setup.md) - Dev environment configuration
- [Testing Guide](testing.md) - Testing procedures

## Resources

- [Pre-commit Documentation](https://pre-commit.com/)
- [golangci-lint Linters](https://golangci-lint.run/usage/linters/)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [mdformat Documentation](https://mdformat.readthedocs.io/)

______________________________________________________________________

**Quick Reference**: Install with `pre-commit install` • Run with
`pre-commit run --all-files` • Update with `pre-commit autoupdate --freeze`
