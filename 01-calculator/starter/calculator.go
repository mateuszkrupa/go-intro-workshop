package main

import (
	"errors"
	"fmt"
)

// ============================================================
// Phase 1: Basic arithmetic (tests provided — make them pass)
// ============================================================

// Add returns the sum of a and b.
func Add(a, b float64) float64 {
	// TODO: implement
	return 0
}

// Sub returns the difference of a and b.
func Sub(a, b float64) float64 {
	// TODO: implement
	return 0
}

// Mul returns the product of a and b.
func Mul(a, b float64) float64 {
	// TODO: implement
	return 0
}

// Div returns the quotient of a divided by b.
// Returns an error if b is zero.
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	// TODO: implement
	return 0, nil
}

// Calculator holds operation history.
type Calculator struct {
	// TODO (Phase 3): add `history []string` field
}

// NewCalculator creates a new Calculator.
func NewCalculator() *Calculator {
	return &Calculator{}
}

// Calculate dispatches an operation by name and returns the result.
// Supported ops: "add", "sub", "mul", "div"
// Returns an error for unknown operators.
//
// TODO (Phase 3): after a successful operation, append to history:
//
//	format: "<a> <op> <b> = <result>"  (e.g. "3 add 5 = 8")
func (c *Calculator) Calculate(op string, a, b float64) (float64, error) {
	switch op {
	case "add":
		return Add(a, b), nil
	case "sub":
		return Sub(a, b), nil
	case "mul":
		return Mul(a, b), nil
	case "div":
		result, err := Div(a, b)
		if err != nil {
			return 0, err
		}
		return result, nil
	default:
		return 0, fmt.Errorf("unknown operator %q", op)
	}
}
