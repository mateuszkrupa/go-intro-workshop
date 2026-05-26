package main

// ============================================================
// Phase 3: History tracking
// ============================================================
//
// The Calculator struct already exists in calculator.go.
// Your job: add a `history []string` field to it, then make
// Calculate() automatically record every successful operation.
//
// Steps:
//   1. Add `history []string` field to Calculator in calculator.go
//   2. In Calculate(), after a successful operation, append an entry
//      Format: "<a> <op> <b> = <result>"  (e.g. "3 add 5 = 8")
//   3. Implement History() and Reset() below

// History returns a copy of the operation history.
func (c *Calculator) History() []string {
	// TODO: implement — return a copy of c.history
	return nil
}

// Reset clears the operation history.
func (c *Calculator) Reset() {
	// TODO: implement — set c.history to empty
}
