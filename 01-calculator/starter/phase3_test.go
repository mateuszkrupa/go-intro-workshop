package main

// ============================================================
// Phase 3 tests (provided — uncomment when you're ready)
// ============================================================
//
// Uncomment the tests below when you start Phase 3.
// Run with: go test -v -run "TestHistory|TestReset" ./...

/*

import "testing"

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
	c.Calculate("div", 5, 0) // error — should NOT be recorded
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

*/
