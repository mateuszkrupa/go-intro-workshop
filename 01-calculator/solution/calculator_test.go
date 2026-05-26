package main

import (
	"errors"
	"math"
	"testing"
)

// ============================================================
// Phase 1 tests
// ============================================================

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{name: "positive numbers", a: 3, b: 5, want: 8},
		{name: "negative numbers", a: -2, b: -3, want: -5},
		{name: "zero", a: 0, b: 5, want: 5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Add(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{name: "positive numbers", a: 10, b: 3, want: 7},
		{name: "negative result", a: 3, b: 10, want: -7},
		{name: "zero", a: 5, b: 0, want: 5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Sub(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Sub(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{name: "positive numbers", a: 4, b: 5, want: 20},
		{name: "negative and positive", a: -3, b: 7, want: -21},
		{name: "multiply by zero", a: 99, b: 0, want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Mul(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Mul(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{name: "valid division", a: 10, b: 2, want: 5, wantErr: false},
		{name: "division with remainder", a: 7, b: 2, want: 3.5, wantErr: false},
		{name: "negative division", a: -9, b: 3, want: -3, wantErr: false},
		{name: "division by zero", a: 5, b: 0, want: 0, wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Div(tc.a, tc.b)

			if tc.wantErr {
				if err == nil {
					t.Errorf("Div(%v, %v) expected error, got nil", tc.a, tc.b)
				}
				return
			}

			if err != nil {
				t.Errorf("Div(%v, %v) unexpected error: %v", tc.a, tc.b, err)
				return
			}

			if math.Abs(got-tc.want) > 1e-9 {
				t.Errorf("Div(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

// ============================================================
// Phase 2 tests
// ============================================================

func TestPow(t *testing.T) {
	tests := []struct {
		name      string
		base, exp float64
		want      float64
	}{
		{name: "2^3", base: 2, exp: 3, want: 8},
		{name: "5^0", base: 5, exp: 0, want: 1},
		{name: "2^-1", base: 2, exp: -1, want: 0.5},
		{name: "3^2", base: 3, exp: 2, want: 9},
		{name: "0^5", base: 0, exp: 5, want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Pow(tc.base, tc.exp)
			if math.Abs(got-tc.want) > 1e-9 {
				t.Errorf("Pow(%v, %v) = %v, want %v", tc.base, tc.exp, got, tc.want)
			}
		})
	}
}

func TestMod(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{name: "10 mod 3", a: 10, b: 3, want: 1, wantErr: false},
		{name: "7 mod 2", a: 7, b: 2, want: 1, wantErr: false},
		{name: "9 mod 3", a: 9, b: 3, want: 0, wantErr: false},
		{name: "negative mod", a: -7, b: 3, want: -1, wantErr: false},
		{name: "mod by zero", a: 5, b: 0, want: 0, wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Mod(tc.a, tc.b)

			if tc.wantErr {
				if err == nil {
					t.Errorf("Mod(%v, %v) expected error, got nil", tc.a, tc.b)
				}
				return
			}

			if err != nil {
				t.Errorf("Mod(%v, %v) unexpected error: %v", tc.a, tc.b, err)
				return
			}

			if math.Abs(got-tc.want) > 1e-9 {
				t.Errorf("Mod(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

// ============================================================
// Phase 3 tests
// ============================================================

func TestHistoryRecordsOperations(t *testing.T) {
	c := NewCalculator()

	c.Calculate("add", 3, 5)
	c.Calculate("mul", 4, 2)
	c.Calculate("div", 10, 3)

	history := c.History()
	if len(history) != 3 {
		t.Fatalf("History() length = %d, want 3", len(history))
	}

	expected := []string{
		"3 add 5 = 8",
		"4 mul 2 = 8",
		"10 div 3 = 3.3333333333333335",
	}
	for i, want := range expected {
		if history[i] != want {
			t.Errorf("History()[%d] = %q, want %q", i, history[i], want)
		}
	}
}

func TestHistorySkipsErrors(t *testing.T) {
	c := NewCalculator()

	c.Calculate("add", 1, 2)
	c.Calculate("div", 5, 0)
	c.Calculate("sub", 9, 4)

	history := c.History()
	if len(history) != 2 {
		t.Fatalf("History() length = %d, want 2 (errors not recorded)", len(history))
	}

	if history[0] != "1 add 2 = 3" {
		t.Errorf("History()[0] = %q, want %q", history[0], "1 add 2 = 3")
	}
	if history[1] != "9 sub 4 = 5" {
		t.Errorf("History()[1] = %q, want %q", history[1], "9 sub 4 = 5")
	}
}

func TestHistoryIsCopy(t *testing.T) {
	c := NewCalculator()

	c.Calculate("add", 1, 1)
	h := c.History()
	h[0] = "tampered"

	if c.History()[0] != "1 add 1 = 2" {
		t.Error("History() did not return a copy — modifying the slice affected internal state")
	}
}

func TestReset(t *testing.T) {
	c := NewCalculator()

	c.Calculate("add", 5, 5)
	c.Calculate("mul", 2, 3)
	c.Reset()

	history := c.History()
	if len(history) != 0 {
		t.Errorf("after Reset(): History() = %v, want empty", history)
	}
}

func TestHistoryEmpty(t *testing.T) {
	c := NewCalculator()

	history := c.History()
	if history == nil {
		t.Error("History() should return empty slice, not nil")
	}
	if len(history) != 0 {
		t.Errorf("History() length = %d, want 0", len(history))
	}
}

// ============================================================
// Phase 4 tests
// ============================================================

func TestMathError(t *testing.T) {
	err := &MathError{Op: "div", A: 5, B: 0, Msg: "division by zero"}

	want := "math: div(5, 0): division by zero"
	if got := err.Error(); got != want {
		t.Errorf("MathError.Error() = %q, want %q", got, want)
	}

	var e error = err
	if e.Error() != want {
		t.Errorf("error interface: got %q, want %q", e.Error(), want)
	}
}

func TestCalculateReturnsCustomError(t *testing.T) {
	c := NewCalculator()

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
