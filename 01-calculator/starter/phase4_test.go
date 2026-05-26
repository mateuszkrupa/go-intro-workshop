package main

// ============================================================
// Phase 4 tests (provided — uncomment when you're ready)
// ============================================================
//
// Uncomment the tests below when you start Phase 4.
// Run with: go test -v -run "TestMathError|TestDivReturnsCustomError|TestCalculateReturnsCustomError" ./...

/*

import (
	"errors"
	"testing"
)

func TestMathError(t *testing.T) {
	err := &MathError{Op: "div", A: 5, B: 0, Msg: "division by zero"}

	want := "math: div(5, 0): division by zero"
	if got := err.Error(); got != want {
		t.Errorf("MathError.Error() = %q, want %q", got, want)
	}

	// Verify it satisfies the error interface
	var e error = err
	if e.Error() != want {
		t.Errorf("error interface: got %q, want %q", e.Error(), want)
	}
}

func TestDivReturnsCustomError(t *testing.T) {
	_, err := Div(5, 0)
	if err == nil {
		t.Fatal("Div(5, 0) should return error")
	}

	var mathErr *MathError
	if !errors.As(err, &mathErr) {
		t.Fatalf("expected *MathError, got %T: %v", err, err)
	}

	if mathErr.Op != "div" {
		t.Errorf("MathError.Op = %q, want %q", mathErr.Op, "div")
	}
	if mathErr.Msg != "division by zero" {
		t.Errorf("MathError.Msg = %q, want %q", mathErr.Msg, "division by zero")
	}
}

func TestCalculateReturnsCustomError(t *testing.T) {
	c := NewCalculator()
	if c == nil {
		t.Fatal("NewCalculator returned nil")
	}

	_, err := c.Calculate("unknown", 1, 2)
	if err == nil {
		t.Fatal("Calculate with unknown op should return error")
	}

	var mathErr *MathError
	if !errors.As(err, &mathErr) {
		t.Fatalf("expected *MathError, got %T: %v", err, err)
	}

	if mathErr.Op != "calculate" {
		t.Errorf("MathError.Op = %q, want %q", mathErr.Op, "calculate")
	}
	if mathErr.A != 1 || mathErr.B != 2 {
		t.Errorf("MathError operands = (%v, %v), want (1, 2)", mathErr.A, mathErr.B)
	}
}

*/
