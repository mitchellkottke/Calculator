/*
Calculator calculates basic math expressions.

Usage:

	calculator <expression>

	expression - A complete math expression. Supported operators are:
					+, -, *, /, ^, ()
*/
package main

import (
	"fmt"
	"os"
	"strings"

	"calculator/calculations"
)

func main() {
	var exp string = parseArgs()

	ans, failed := calculations.Evaluate(exp)
	if !failed {
		fmt.Println("Answer is:", ans)
	}
}

func parseArgs() string {
	if len(os.Args) != 2 || len(os.Args[1]) == 0 {
		usage(os.Args[0])
	}

	return strings.Join(os.Args[1:], "")
}

func usage(progName string) {
	fmt.Println("Usage: ", progName, "<expression>")
	fmt.Println("\texpression - The expression to evaluate")
	fmt.Println("\t\tSupported operators are: +, -, *, /, ^, ()")
	os.Exit(1)
}
