package main

// ============================================================
// Phase 4: Calculate dispatcher & custom errors
// ============================================================

// MathError represents a math operation failure with context.
type MathError struct {
	Op   string  // the operation name, e.g. "div", "mod", "calculate"
	A, B float64 // the operands
	Msg  string  // what went wrong
}

// Error returns a formatted error string.
// Format: "math: <op>(<a>, <b>): <msg>"
// Example: "math: div(5, 0): division by zero"
func (e *MathError) Error() string {
	// TODO: implement
	return ""
}


