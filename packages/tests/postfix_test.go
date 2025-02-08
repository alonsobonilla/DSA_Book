package test

import (
	"exercises/cmd/infixToPostfix"
	"fmt"
	"testing"
	"unicode"
)

func TestComputePostFixExpression(t *testing.T) {
	infixExpressions := map[string]float64{
		"2+2-3+2":         3,
		"2^2^3-1":         255,
		"(2^2)^3+6/2":     67,
		"2-2-2+3":         1,
		"2-5-9+2-12/2^2":  -13,
		"9/3+2-3^3":       -22,
		"100/2":           50,
		"(100+200/10)/10": 12,
	}
	for infixString, expected := range infixExpressions {
		var infixArr []string

		num := ""
		for _, e := range infixString {
			if unicode.IsDigit(e) {
				num += string(e)
			} else {
				if num != "" {
					infixArr = append(infixArr, num)
					num = ""
				}
				infixArr = append(infixArr, string(e))
			}
		}
		infixArr = append(infixArr, num)

		fmt.Print(infixArr)
		fmt.Print("Length: ", len(infixArr))
		postfix := infixToPostfix.InfixToPostfix(infixArr)
		result := infixToPostfix.ComputePostFixExpression(postfix)
		if result != expected {
			t.Fatalf("The infix expression (%s) is became to postfix: %s is expected %f but was obtained %f", infixString, postfix, expected, result)
		}
	}
}
