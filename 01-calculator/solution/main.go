package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	calc := NewCalculator()

	// If args provided, run once and exit: go run . 3 add 5
	if len(os.Args) == 4 {
		run(calc, os.Args[1], os.Args[2], os.Args[3])
		return
	}

	// Otherwise start interactive REPL
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Calculator — type: <a> <op> <b> (e.g. 3 add 5)")
	fmt.Println("Operators: add, sub, mul, div, pow, mod")
	fmt.Println("Commands: history, reset, quit")
	fmt.Println()

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if line == "quit" || line == "exit" {
			break
		}
		if line == "history" {
			for _, entry := range calc.History() {
				fmt.Println(" ", entry)
			}
			continue
		}
		if line == "reset" {
			calc.Reset()
			fmt.Println("History cleared.")
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 3 {
			fmt.Println("Usage: <a> <op> <b>")
			continue
		}

		run(calc, parts[0], parts[1], parts[2])
	}
}

func run(calc *Calculator, aStr, op, bStr string) {
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		fmt.Printf("Error: invalid number %q\n", aStr)
		return
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		fmt.Printf("Error: invalid number %q\n", bStr)
		return
	}

	result, err := calc.Calculate(op, a, b)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(result)
}
