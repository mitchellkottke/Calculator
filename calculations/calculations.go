// Package calculations is a utility to evaluate basic string math expressions with
// order of operations.
// Supported operators:
//
//	Addition: +
//	Subtraction: -
//	Multiplication: *
//	Division: /
//	Exponents: ^
//	Grouping: ( )
//
// Functions for use:
//
//	Evaluate(string)
package calculations

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"calculator/stack"
)

// This works like an enum in C
// Ranking of operators
const (
	LPAREN = iota
	ADD_SUB
	MULT_DIV
	EXP
	RPAREN
)

// Evaluate will evaluate the given math expression
// @param exp - Expression to evaluate, expects a valid expression with no spaces
// @ret ans - The answer on success
// @ret err - Whether there was an error, if there was ans is garbage
func Evaluate(exp string) (ans float64, failed bool) {
	answer := stack.New()
	if !parseExp(exp, answer) {
		fmt.Println("Failed to parse expression")
		return 0.0, true
	}

	return answer.Pop().(float64), false
}

func popStackFloat(stack *stack.Stack) (val float64, err bool) {
	value := stack.Pop()
	if value == nil {
		return 0.0, true
	}
	return value.(float64), false
}

// Uses shunting yard https://en.wikipedia.org/wiki/Shunting_yard_algorithm
func parseExp(exp string, operands *stack.Stack) bool {
	num_start_ind := -1
	found_num := false
	found_dot := false
	chars := strings.Split(exp, "")
	operators := stack.New()

	for i, c := range chars {
		if isNumber(c) {
			if !found_num {
				num_start_ind = i
				found_num = true
			}
		} else if isDot(c) {
			if found_dot {
				fmt.Println("Found multiple decimal points in one number")
				return false
			}
			if !found_num {
				found_num = true
				num_start_ind = i
			}
			found_dot = true
		} else if isOper(c) {
			if found_num {
				if !parseAndAddNum(operands, exp, num_start_ind, i, found_num, found_dot) {
					return false
				}
				found_num = false
				if found_dot {
					found_dot = false
				}
			}
			if !handleOper(operands, operators, c) {
				return false
			}
		} else {
			fmt.Println("Found unsupported symbol in expression:", c)
			return false
		}
	}

	if found_num || found_dot {
		if !parseAndAddNum(operands, exp, num_start_ind, len(exp), found_num, found_dot) {
			return false
		}
	}

	//Handle remaining operations
	for oper := operators.Peek(); oper != nil; oper = operators.Peek() {
		operators.Pop()

		num2, err2 := popStackFloat(operands)
		num1, err1 := popStackFloat(operands)
		if err1 || err2 {
			fmt.Println("Tried to pop empty stack")
			return false
		}

		if !executeOper(num1, num2, oper.(string), operands) {
			return false
		}
	}

	return true
}

func parseAndAddNum(stack *stack.Stack, exp string, start_ind int, end_ind int, found_num bool, found_dot bool) bool {
	len := end_ind - start_ind

	if !found_num || len <= 0 {
		fmt.Println("Internal error, found number of 0 length at position:", start_ind)
		return false
	}

	if found_dot && len == 1 {
		fmt.Println("Error decimal point needs a number attached at position:", start_ind)
		return false
	}

	num, err := strconv.ParseFloat(exp[start_ind:end_ind], 64)

	if err != nil {
		fmt.Println(err)
		return false
	}

	stack.Push(num)

	return true
}

func handleOper(operands *stack.Stack, operators *stack.Stack, oper string) bool {
	if oper[0] == '(' || oper[0] == ')' {
		return handleParen(operands, operators, oper)
	}

	if operators.Peek() == nil || operComp(oper, operators.Peek().(string)) > 0 {
		operators.Push(oper)
		return true
	}
	//Top of operators stack is a higher or equal rank to this operator
	for curr_oper := operators.Peek(); curr_oper != nil && operComp(oper, curr_oper.(string)) <= 0; curr_oper = operators.Peek() {
		//Time to execute this operator
		operators.Pop()

		num2, err2 := popStackFloat(operands)
		num1, err1 := popStackFloat(operands)
		if err1 || err2 {
			fmt.Println("Tried to pop empty stack")
			return false
		}

		if !executeOper(num1, num2, curr_oper.(string), operands) {
			return false
		}
	}

	operators.Push(oper)
	return true
}

func handleParen(operands *stack.Stack, operators *stack.Stack, oper string) bool {
	var curr_oper interface{}

	if oper[0] == '(' {
		operators.Push(oper)
		return true
	}

	for curr_oper = operators.Peek(); curr_oper != nil && curr_oper.(string)[0] != '('; curr_oper = operators.Peek() {
		operators.Pop()

		num2, err2 := popStackFloat(operands)
		num1, err1 := popStackFloat(operands)
		if err1 || err2 {
			fmt.Println("Tried to pop empty stack")
			return false
		}

		if !executeOper(num1, num2, curr_oper.(string), operands) {
			return false
		}
	}

	if curr_oper == nil {
		fmt.Println("Error no closing parenthesis found")
		return false
	}

	operators.Pop() //For the (

	return true
}

func executeOper(num1 float64, num2 float64, oper string, stack *stack.Stack) bool {
	switch oper[0] {
	case '+':
		stack.Push(num1 + num2)
	case '-':
		stack.Push(num1 - num2)
	case '*':
		stack.Push(num1 * num2)
	case '/':
		if num2 == 0 {
			fmt.Println("Error cannot divide by zero")
			return false
		}
		stack.Push(num1 / num2)
	case '^':
		stack.Push(math.Pow(num1, num2))
	default:
		fmt.Println("Error invalid operator found:", oper)
		return false
	}

	return true
}

// This is guarenteed to get valid operators because of the isOper checks
func operComp(o1 string, o2 string) int {
	return assignOperRank(o1) - assignOperRank(o2)
}

func assignOperRank(oper string) int {
	switch oper[0] {
	case '(':
		return LPAREN
	case '+', '-':
		return ADD_SUB
	case '*', '/':
		return MULT_DIV
	case '^':
		return EXP
	case ')':
		return RPAREN
	default:
		fmt.Println("Error unexpected operator in assignOperRank:", oper)
		fmt.Println("This should be impossible")
		return -1
	}
}

func isNumber(c string) bool {
	return c[0] >= '0' && c[0] <= '9'
}

func isDot(c string) bool {
	return c[0] == '.'
}

func isOper(c string) bool {
	return c[0] == '+' || c[0] == '-' || c[0] == '*' || c[0] == '/' || c[0] == '^' || c[0] == '(' || c[0] == ')'
}
