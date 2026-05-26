# Project 1: Calculator

Build a calculator package in Go across four phases — from basic functions to custom error types.

## How to Build & Run

```bash
# Run the interactive calculator
go run .

# Or run a single calculation and exit
go run . 3 add 5

# Run all tests
go test ./...

# Run tests with verbose output (see each test name)
go test -v ./...

# Run only specific tests (e.g. Phase 1)
go test -v -run "TestAdd|TestSub|TestMul|TestDiv" ./...
```

> **Windows note:** All `go` commands work the same in PowerShell and Command Prompt. Use PowerShell for the best experience.

The interactive calculator reads expressions like `3 add 5` and prints the result. Type `quit` to exit.
You can also pass arguments directly: `go run . 10 div 2` prints the result and exits.

---

## Phase 1: Basic Arithmetic (tests provided)

Implement four arithmetic functions in `calculator.go` and make the pre-written tests pass.

### Tasks

1. **Run the tests first and watch them fail:**
   ```bash
   go test -v -run "TestAdd|TestSub|TestMul|TestDiv" ./...
   ```
   Read the output — it tells you exactly what each function should return.
2. Open `calculator.go` and implement the `Add` function — return the sum of `a` and `b`.
3. Implement the `Sub` function — return the difference of `a` and `b`.
4. Implement the `Mul` function — return the product of `a` and `b`.
5. Implement the `Div` function — return the quotient of `a / b`, or an error if `b` is zero.
6. Run the tests again and confirm Phase 1 tests pass:
   ```bash
   go test -v -run "TestAdd|TestSub|TestMul|TestDiv" ./...
   ```
7. Try the interactive calculator:
   ```bash
   go run .
   > 3 add 5
   > 10 div 0
   > quit
   ```

### Hints

- Go functions can return multiple values: `func Div(a, b float64) (float64, error)`
- Use `errors.New("message")` from the `errors` package to create error values.
- You need to import the `"errors"` package: `import "errors"` (already done in `calculator.go`).
- The tests use table-driven style with `t.Run` — read through `calculator_test.go` to understand what's expected.

---

## Phase 2: Pow & Mod (you write the tests!)

Implement `Pow` and `Mod` in `calculator.go`, then **write your own tests** in `calculator_test.go`.

### Tasks

1. Implement `Pow(base, exp float64) float64` — hint: use `math.Pow`.
2. Implement `Mod(a, b float64) (float64, error)` — return error if `b` is zero, hint: use `math.Mod`.
3. Write `TestPow` in `calculator_test.go` using the table-driven pattern. Test at least:
   - `2^3 = 8`
   - `5^0 = 1`
   - `2^-1 = 0.5`
4. Write `TestMod` in `calculator_test.go`. Test at least:
   - `10 mod 3 = 1`
   - `7 mod 2 = 1`
   - mod by zero returns an error
5. Run your new tests:
   ```bash
   go test -v -run "TestPow|TestMod" ./...
   ```

### Hints

- You'll need to add `"math"` to your imports: `import "math"` — then use `math.Pow` and `math.Mod`.
- For `Mod`, you also need `"errors"` for the error case (same as `Div`).
- Follow the exact same table-driven pattern as `TestDiv` — copy it and adapt.
- For `Mod`, the error case is the same pattern as `Div` (division by zero).

---

## Phase 3: History Tracking (tests provided)

Make the calculator automatically record every successful operation.

### Tasks

1. **Uncomment the tests in `phase3_test.go`, then run them to see what's expected:**
   ```bash
   go test -v -run "TestHistory|TestReset" ./...
   ```
2. Add a `history []string` field to the `Calculator` struct in `calculator.go`.
3. In `Calculate()`, after a successful operation, append a history entry.
   - Format: `"<a> <op> <b> = <result>"` — e.g. `"3 add 5 = 8"`
   - Don't record failed operations (errors).
4. Implement `History()` in `phase3.go` — return a **copy** of the history slice.
5. Implement `Reset()` in `phase3.go` — clear the history.
6. Run the tests:
   ```bash
   go test -v -run "TestHistory|TestReset" ./...
   ```
7. Try it interactively — your history builds up:
   ```bash
   go run .
   > 3 add 5
   > 10 div 2
   > quit
   ```

### Hints

- Use `fmt.Sprintf("%g %s %g = %g", a, op, b, result)` to format history entries.
- `History()` should return a copy: `result := make([]string, len(c.history))` + `copy(result, c.history)`.
- Only append to history when `err == nil` — failed operations should not be recorded.
- You'll need to restructure `Calculate` slightly so you have the `result` value before returning it.

---

## Phase 4: Custom errors (optional)

Upgrade your error handling to use a custom `MathError` type instead of plain strings.

### Tasks

1. **Uncomment the tests in `phase4_test.go`, then run them:**
   ```bash
   go test -v -run "TestMathError|TestDivReturnsCustomError|TestCalculateReturnsCustomError" ./...
   ```
2. Look at the `MathError` struct in `phase4.go` — the fields are already defined. Implement the `Error() string` method so it satisfies the `error` interface.
   - Format: `"math: <op>(<a>, <b>): <msg>"` — e.g. `"math: div(5, 0): division by zero"`
3. Update `Div` in `calculator.go` to return `&MathError{Op: "div", ...}` instead of `errors.New(...)`.
4. Update `Mod` in `phase2.go` to return `&MathError{Op: "mod", ...}` instead of `errors.New(...)`.
5. Update `Calculate` in `phase3.go` to return `&MathError{Op: "calculate", ...}` for unknown operators instead of `fmt.Errorf(...)`.
6. Run all tests:
   ```bash
   go test -v ./...
   ```

### Hints

- You'll need `"fmt"` for `fmt.Sprintf` in the `Error()` method.
- A custom error type just needs an `Error() string` method:
  ```go
  func (e *MathError) Error() string {
      return fmt.Sprintf("math: %s(%g, %g): %s", e.Op, e.A, e.B, e.Msg)
  }
  ```
- In tests, use `errors.As` (from the `"errors"` package) to check the error type:
  ```go
  var mathErr *MathError
  if errors.As(err, &mathErr) {
      // check mathErr.Op, mathErr.Msg, etc.
  }
  ```
- You don't need to change `main.go` — it already prints `err` which will use your `Error()` method.

---

## Run & Test Commands

| Command | Purpose |
|---------|---------|
| `go test ./...` | Run all tests (all phases) |
| `go test -v ./...` | Verbose output |
| `go test -cover ./...` | Coverage report |
| `go test -run TestAdd ./...` | Run a specific test |
| `go run .` | Start the interactive calculator |

## Stretch Goals

These are optional challenges for when you finish all four phases:

1. **Undo** — Add an `Undo()` method to `Calculator` that reverts the last operation (hint: store previous results too).
2. **Expression evaluator** — Add `Eval(expr string) (float64, error)` that parses `"3 + 5 * 2"` (left-to-right, no precedence).
3. **Error wrapping** — Use `fmt.Errorf("calculate: %w", err)` to wrap errors and test with `errors.Is` / `errors.Unwrap`.

---

## Phase 5: Symbol operators (optional)

Support `+`, `-`, `*`, `/`, `^`, `%` as aliases for `add`, `sub`, `mul`, `div`, `pow`, `mod`.

### Tasks

1. Update `Calculate` to recognize both word and symbol forms:
   - `"+"` → add, `"-"` → sub, `"*"` → mul, `"/"` → div, `"^"` → pow, `"%"` → mod
2. Write tests for the symbol operators (e.g. `c.Calculate("+", 3, 5)` returns 8).
3. Try it in the REPL:
   ```bash
   > 3 + 5
   > 10 / 2
   > 2 ^ 8
   ```

### Hints

- You can normalize the operator at the top of `Calculate` before the switch:
  ```go
  switch op {
  case "+", "add":
      // ...
  }
  ```
- Or create a helper: `func normalizeOp(op string) string` that maps symbols to words.

---

## Phase 6: Input validation (optional)

Add a `ValidateOperator` function that checks the operator against a predefined slice of supported operators.

### Tasks

1. Define a package-level slice of valid operators:
   ```go
   var supportedOps = []string{"add", "sub", "mul", "div", "pow", "mod"}
   ```
2. Create a `ValidateOperator(op string) error` function that iterates over `supportedOps` and returns an error if `op` is not found.
3. Call `ValidateOperator` at the start of `Calculate` — return early if validation fails.
4. Write tests: valid operators return nil, unknown operators return an error.

### Hints

- Loop over the slice with `for _, valid := range supportedOps`.
- If you did Phase 5, add the symbol forms to the slice too: `"+", "-", "*", "/", "^", "%"`.
- This teaches working with slices, iteration, and separating validation from logic.

---

## Phase 7: Operator as a type (optional)

Define operators as a custom type with methods — teaches Go's type system and `iota`.

### Tasks

1. Define a type: `type Operator int`
2. Define constants using `iota`:
   ```go
   const (
       OpAdd Operator = iota
       OpSub
       OpMul
       OpDiv
       OpPow
       OpMod
   )
   ```
3. Add a `String() string` method on `Operator` (satisfies `fmt.Stringer`).
4. Add a `ParseOperator(s string) (Operator, error)` function that converts strings like `"add"`, `"+"` to the enum.
5. Refactor `Calculate` to accept `Operator` instead of `string`.
6. Update `main.go` to parse the operator string into an `Operator` before calling `Calculate`.
7. Write tests for `ParseOperator` and `Operator.String()`.

### Hints

- The `String()` method lets you print operators nicely: `fmt.Println(OpAdd)` → `"add"`.
- `ParseOperator` is the reverse — it maps user input to the typed constant.
- This pattern is common in Go for representing a fixed set of values (no built-in enums).
