package test

import (
	"exercises/cmd/infixToPostfix"
	"testing"
)

func TestComputePostFixExpression(t *testing.T) {
	infixExpressions := map[string]float64{
		"2+2-3+2":       3,
		"2^2^3-1":       255,
		"(2^2)^3+6/2":   67,
		"2-2-2+3":       1,
		"2-5-9+2-8/2^2": -12,
		"9/3+2-3^3":     -22,
	}
	for infix, expected := range infixExpressions {
		postfix := infixToPostfix.InfixToPostfix(infix)
		result := infixToPostfix.ComputePostFixExpression(postfix)
		if result != expected {
			t.Fatalf("The infix expression (%s) is became to postfix: %s is expected %f but was obtained %f", infix, string(postfix), expected, result)
		}
	}
}
