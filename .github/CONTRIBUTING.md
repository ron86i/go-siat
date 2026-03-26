# Contribution Guide

Thank you for your interest in contributing to **go-siat**. This document describes the process and best practices.

## 📋 Prerequisites

- Go 1.25 or higher
- Git
- Familiarity with SIAT (Integrated Tax Administration System)
- Read [ARCHITECTURE.md](ARCHITECTURE.md)

---

## 🚀 Getting Started

### 1. Fork and Clone

```bash
# Fork the repository on GitHub
git clone https://github.com/your-user/go-siat.git
cd go-siat
```

### 2. Create a Branch

```bash
# Create a branch for your feature or fix
git checkout -b feature/my-feature
# or
git checkout -b fix/issue-number
```

### 3. Setup the Environment

```bash
# Download dependencies
go mod tidy

# Verify everything compiles
go build ./...

# Run the tests
go test ./...
```

---

## 💻 Development

### Code Standards

#### Comments and Documentation

- **Godoc**: Every public type, function, or method must have a comment in godoc format.
- **Examples**: Include examples in tests for documentation.

```go
// SolicitudCuis obtains the Unique System Initiation Code required to operate with SIAT.
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - config: Configuration with authentication token
//   - req: CUIS request built with builder
//
// Returns EnvelopeResponse with the SOAP response or error.
func (s *SiatCodigosService) SolicitudCuis(ctx context.Context, config Config, req models.Cuis) (*soap.EnvelopeResponse[codigos.CuisResponse], error) {
```

#### Naming

- Variables: `camelCase` (e.g., `cufsResponse`, `errorMessage`)
- Constants: `UPPER_SNAKE_CASE` (e.g., `DEFAULT_TIMEOUT`, `MAX_RETRIES`)
- Public Functions: `PascalCase` (e.g., `SolicitudCuis`)
- Private Methods: `camelCase` (e.g., `buildSoapEnvelope`)

#### Errors

- Always use `fmt.Errorf` or `errors.New` with context.
- Avoid returning `nil` for errors without using `error` as the type.

```go
// Good
if err != nil {
    return nil, fmt.Errorf("failed to compile SOAP: %w", err)
}

// Avoid
if err != nil {
    return nil, err  // Context is lost
}
```

#### Type-Safety

- Prefer specific types over `interface{}`.
- Use builders when possible.

```go
// Good
type SolicitudCuis struct {
    CodigoAmbiente int64
    NIT            int64
}

// Avoid
type SolicitudCuis map[string]interface{}
```

### Test Structure

Tests should follow the table-driven test pattern:

```go
func TestSolicitudCuis(t *testing.T) {
    tests := []struct {
        name      string
        req       models.Cuis
        wantErr   bool
        setupMock func()
    }{
        {
            name: "valid request",
            req: models.Codigos().NewCuisBuilder().
                WithCodigoAmbiente(1).
                Build(),
            wantErr: false,
        },
        {
            name: "invalid environment",
            req: models.Codigos().NewCuisBuilder().
                WithCodigoAmbiente(99).
                Build(),
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### Commits

- **Clear messages**: Describe *what* and *why*, not *how*.
- **Scope**: `[feature|fix|docs|test|refactor]: description`

```bash
# Good
git commit -m "fix: validate NIT before sending CUFD" \
           -m "SIAT rejects CUFD with inconsistencies in the NIT" \
           -m "Closes #123"

# Avoid
git commit -m "bug fix"
```

---

## 🧪 Testing

### Minimum Coverage

- **New code**: 80% coverage.
- **Changes to existing code**: Do not decrease the current coverage.

```bash
# Run tests with coverage
go test -cover ./...

# Generate HTML coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Test Types

| Type | Location | Purpose |
|------|-----------|-----------|
| **Unit** | `*_test.go` in the same package | Verify logic in isolation |
| **Integration** | `*_integration_test.go` | Verify behavior with real SIAT (CI only) |
| **Regression** | `testdata/` | Validate against known responses |

### Environment Variables for Tests

```bash
# Integration tests require these values
export SIAT_API_TOKEN="your_token"
export SIAT_NIT="your_nit"
export SIAT_BASE_URL="https://pilotosiatservicios.impuestos.gob.bo/v2"

# Run integration tests
go test -tags=integration ./...
```

---

## 📦 Contribution Types

### 1. New Features

- Create an issue first describing the change.
- Wait for approval before starting.
- Include tests and documentation.
- Update examples if applicable.

### 2. Bug Fixes

- Include a test that reproduces the bug.
- Fix it.
- Verify that the test passes.
- Document it in the commit.

### 3. Documentation Improvements

- Update `README.md`, `ARCHITECTURE.md`, or create new `.md` files.
- Improve godoc comments.
- Add examples.
- Does not need prior approval.

### 4. Performance Improvements

- Include benchmarks demonstrating the improvement.
- Do not sacrifice readability.
- Document the design decision.

---

## 🔄 Pull Request Process

### Step 1: Prepare the PR

```bash
# Verify everything is ok
go fmt ./...
go vet ./...
go test ./...

# Push to your branch
git push origin feature/my-feature
```

### Step 2: Create the Pull Request

On GitHub:

1. Title: `[type] brief description`
2. Description: Explain what changed, why, and what tests were included.
3. Reference: `Closes #issueNumber` if applicable.

**Description Template:**

```markdown
## Description
Brief description of the change.

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation improvement

## Testing
- [ ] Unit tests added/modified
- [ ] Integration tests (if applicable)
- [ ] Coverage >= 80%

## Checklist
- [ ] Code follows project standards
- [ ] All tests pass
- [ ] Documentation updated
- [ ] No linting issues
```

### Step 3: Review

- Wait for review from maintainers.
- Respond to comments in the PR.
- Make requested changes.
- Re-request review.

### Step 4: Merge

Once approved, a maintainer will perform the merge.

---

## 📋 Quality Checklist

Before opening a PR:

- [ ] `go fmt ./...` has been executed.
- [ ] `go vet ./...` no errors.
- [ ] `go test ./...` all tests pass.
- [ ] Test coverage >= 80% for new code.
- [ ] All public types/functions have godoc documentation.
- [ ] Examples in comments if complex.
- [ ] No unrelated changes (keep PRs focused).
- [ ] Commits with clear messages.
- [ ] `README.md` updated if there are API changes.

---

## 🐛 Reporting Bugs

1. Verify that a similar issue does not already exist.
2. Create an issue with:
   - Descriptive title.
   - Detailed description.
   - Code that reproduces the problem.
   - Go and go-siat versions.
   - SIAT version (production/testing).

**Example:**

```markdown
## Description
CUFD rejected with error "Invalid Control Code"

## Steps to Reproduce
1. Call SolicitudCufd with...
2. Observe the error...

## Expected Output
Valid CUFD code

## System Information
- Go: 1.25
- go-siat: v1.2.0
- SIAT: Production
```

---

## 🎓 Project Conventions

### Package Structure

- `pkg/`: Public API.
- `internal/`: Private implementation.
- Keep packages focused on a single responsibility.

### Dependencies

- Keep `go.mod` clean.
- Justify new external dependencies.
- Prefer stdlib when possible.

### Semantic Versioning

- MAJOR: Breaking changes.
- MINOR: New features (backward compatible).
- PATCH: Bug fixes.

---

## 💬 Communication

- **Issues**: To report bugs or propose features.
- **Discussions**: For general questions and conversations.
- **Email**: Contact maintainers if necessary.

---

## 📖 Resources

- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [ARCHITECTURE.md](ARCHITECTURE.md) - Project Design
- [SIAT Documentation](https://www.impuestos.gob.bo/) - Official SIAT Documentation

---

## ✨ Thanks

Thank you for contributing to **go-siat**. Your contributions make electronic invoicing in Bolivia more accessible for all developers.
