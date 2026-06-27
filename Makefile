.PHONY: help build test lint coverage clean

help:
	@echo "Available targets:"
	@echo "  build       - Build the application"
	@echo "  test        - Run tests"
	@echo "  coverage    - Run tests with coverage report"
	@echo "  lint        - Run golangci-lint"
	@echo "  clean       - Clean build artifacts"

build:
	@echo "Building application..."
	@go build -o bin/app ./cmd/app

test:
	@echo "Running tests..."
	@go test -v ./...

coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
	@go tool cover -func=coverage.out

lint:
	@echo "Running linter..."
	@golangci-lint run ./...

clean:
	@echo "Cleaning..."
	@rm -f bin/app coverage.out coverage.html
	@go clean
