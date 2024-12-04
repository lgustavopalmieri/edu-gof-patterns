//Explanation:
//Expression Interface: Declares the Interpret method, which all concrete expressions must implement.
//Terminal & Non-Terminal Expressions: NumberExpression represents numbers, which are terminal expressions. AddExpression and SubtractExpression represent operations on expressions, serving as non-terminal expressions.
//Interpreter Logic: Each expression's Interpret method recursively evaluates the expression based on its specific logic (e.g., addition or subtraction).
//Parsing Mechanism: Demonstrates a simple parsing scheme where tokens are interpreted to generate an expression tree, which is then evaluated.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Expression interface defines the interpret method
type Expression interface {
	Interpret() int
}

// NumberExpression represents a terminal numeric expression
type NumberExpression struct {
	value int
}

// Interpret returns the value of the number
func (n *NumberExpression) Interpret() int {
	return n.value
}

// AddExpression represents a '+' operation between two expressions
type AddExpression struct {
	left, right Expression
}

// Interpret evaluates the addition of two expressions
func (a *AddExpression) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}

// SubtractExpression represents a '-' operation between two expressions
type SubtractExpression struct {
	left, right Expression
}

// Interpret evaluates the subtraction of two expressions
func (s *SubtractExpression) Interpret() int {
	return s.left.Interpret() - s.right.Interpret()
}

// ParseToken parses a token and creates the appropriate expression
func ParseToken(context []string) Expression {
	if len(context) == 0 {
		return nil
	}

	// Process numbers as terminal expressions
	if value, err := strconv.Atoi(context[0]); err == nil {
		return &NumberExpression{value: value}
	}

	// Interpret the rest as operators and expressions
	switch context[0] {
	case "+":
		left := ParseToken(context[1:])
		right := ParseToken(context[2:])
		return &AddExpression{left: left, right: right}
	case "-":
		left := ParseToken(context[1:])
		right := ParseToken(context[2:])
		return &SubtractExpression{left: left, right: right}
	}

	return nil
}

func main() {
	context := strings.Split("3 + 5 - 2", " ")
	// This requires a better parsing mechanism to handle proper order.
	// For simplicity, we consider each operator is followed by correct operands.
	// The emphasis is on demonstrating expression evaluation.

	expression := ParseToken(context)
	result := expression.Interpret()
	fmt.Printf("Result of expression: %d\n", result)
}

//OUTPUT
//
//Result of expression: 3
