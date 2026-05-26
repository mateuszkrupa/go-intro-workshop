package main

import (
	"math"
	"testing"
)

// ============================================================
// Phase 1 tests (provided — make these pass first)
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
