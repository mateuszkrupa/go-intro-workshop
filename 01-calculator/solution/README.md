# Project 1: Calculator — Solution

This is the complete reference implementation of the calculator package. It implements all four arithmetic functions (`Add`, `Sub`, `Mul`, `Div`) with proper error handling for division by zero.

## What's Included

- `calculator.go` — Full implementation of `Add`, `Sub`, `Mul`, and `Div`
- `calculator_test.go` — Table-driven tests covering all functions and edge cases
- `cmd/calc/main.go` — CLI tool that dispatches to calculator functions

## Run Tests

```bash
go test ./...
```

All tests should pass with zero failures.

## Verify Coverage

```bash
go test -cover ./...
```

Coverage should be at least 90%.

## Run the CLI

```bash
go run ./cmd/calc add 3 5
go run ./cmd/calc div 10 2
go run ./cmd/calc div 10 0
```
