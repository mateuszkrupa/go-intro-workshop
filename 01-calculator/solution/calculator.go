package main

import (
	"fmt"
	"math"
)

// ============================================================
// Phase 1: Basic arithmetic
// ============================================================

// Add returns the sum of a and b.
func Add(a, b float64) float64 {
	return a + b
}

// Sub returns the difference of a and b.
func Sub(a, b float64) float64 {
	return a - b
}

// Mul returns the product of a and b.
func Mul(a, b float64) float64 {
	return a * b
}

// Div returns the quotient of a divided by b.
// Returns an error if b is zero.
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &MathError{Op: "div", A: a, B: b, Msg: "division by zero"}
	}
	return a / b, nil
}

// ============================================================
// Phase 2: Pow & Mod
// ============================================================

// Pow returns base raised to the power of exp.
func Pow(base, exp float64) float64 {
	return math.Pow(base, exp)
}

// Mod returns the remainder of a divided by b.
// Returns an error if b is zero.
func Mod(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &MathError{Op: "mod", A: a, B: b, Msg: "modulo by zero"}
	}
	return math.Mod(a, b), nil
}

// ============================================================
// Phase 3: Calculator with history
// ============================================================

// Calculator holds operation history.
type Calculator struct {
	history []string
}

// NewCalculator creates a new Calculator.
func NewCalculator() *Calculator {
	return &Calculator{
		history: []string{},
	}
}

// Calculate dispatches an operation by name and returns the result.
// Records successful operations in history.
// Format: "<a> <op> <b> = <result>"
func (c *Calculator) Calculate(op string, a, b float64) (float64, error) {
	var result float64
	var err error

	switch op {
	case "add":
		result = Add(a, b)
	case "sub":
		result = Sub(a, b)
	case "mul":
		result = Mul(a, b)
	case "div":
		result, err = Div(a, b)
	case "pow":
		result = Pow(a, b)
	case "mod":
		result, err = Mod(a, b)
	default:
		return 0, &MathError{Op: "calculate", A: a, B: b, Msg: "unknown operator: " + op}
	}

	if err != nil {
		return 0, err
	}

	c.history = append(c.history, fmt.Sprintf("%g %s %g = %g", a, op, b, result))
	return result, nil
}

// History returns a copy of the operation history.
func (c *Calculator) History() []string {
	result := make([]string, len(c.history))
	copy(result, c.history)
	return result
}

// Reset clears the operation history.
func (c *Calculator) Reset() {
	c.history = []string{}
}

// ============================================================
// Phase 4: Custom errors
// ============================================================

// MathError represents a math operation failure with context.
type MathError struct {
	Op   string
	A, B float64
	Msg  string
}

// Error returns a formatted error string.
func (e *MathError) Error() string {
	return fmt.Sprintf("math: %s(%g, %g): %s", e.Op, e.A, e.B, e.Msg)
}
