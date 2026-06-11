# golang-samples

A modern Go project template demonstrating best practices and contemporary design patterns in Go development.

## Project Structure

```
.
├── cmd/
│   └── app/                    # Application entry point
│       └── main.go
├── internal/
│   └── calculator/             # Business logic (internal package)
│       ├── calculator.go
│       └── calculator_test.go
├── .github/
│   └── workflows/
│       └── ci.yml             # CI/CD configuration
├── .golangci.yml              # Linter configuration
├── go.mod                     # Module definition
├── Makefile                   # Build and test commands
└── README.md                  # This file
```

## Features

- **Modern Go Patterns**: Demonstrates service pattern with dependency injection
- **Comprehensive Testing**: Unit tests with table-driven test cases
- **Code Coverage**: Coverage reporting integrated into CI/CD pipeline
- **CI/CD Pipeline**: Automated linting, building, and testing on every push and PR
- **Professional Tooling**: Uses golangci-lint and codecov for quality assurance

## Prerequisites

- Go 1.21 or later
- Make (optional, for using Makefile commands)
- golangci-lint (for local linting)

## Getting Started

### Build the Application

```bash
make build
# or
go build -o bin/app ./cmd/app
```

### Run the Application

```bash
./bin/app
```

### Run Tests

```bash
make test
# or
go test -v ./...
```

### Generate Coverage Report

```bash
make coverage
# Generates coverage.html for detailed report
```

### Run Linter

```bash
make lint
# or
golangci-lint run ./...
```

## Modern Design Patterns Used

### 1. **Service Pattern**
The `Calculator` struct encapsulates related operations, demonstrating the service pattern commonly used in Go applications.

```go
calc := calculator.NewCalculator()
result := calc.Add(5, 3)
```

### 2. **Error Handling**
Explicit error handling following Go conventions. Methods that can fail return `(result, error)`.

```go
result, err := calc.Divide(20, 4)
if err != nil {
    log.Fatal(err)
}
```

### 3. **Table-Driven Tests**
Comprehensive test coverage using table-driven test approach for maintainability and readability.

```go
tests := []struct {
    name     string
    a        int
    b        int
    expected int
}{
    {"positive numbers", 5, 3, 8},
    {"negative numbers", -5, -3, -8},
}
```

### 4. **Package Organization**
- `cmd/app`: Application-specific code
- `internal/`: Business logic not meant to be imported by external packages

## CI/CD Workflow

The project includes a comprehensive GitHub Actions workflow that:

1. **Lints** the code using golangci-lint
2. **Builds** the application
3. **Runs tests** and generates coverage reports
4. **Uploads** coverage to codecov

The workflow is triggered on:
- Push to `main` or `develop` branches
- Pull requests to `main` or `develop` branches

## Contributing

When adding new features:

1. Follow the existing project structure
2. Write table-driven tests for new functionality
3. Ensure all tests pass: `make test`
4. Run linter locally: `make lint`
5. Generate coverage report: `make coverage`

## License

See LICENSE file for details.
