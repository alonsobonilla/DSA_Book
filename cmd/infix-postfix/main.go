package main

import (
	stack "exercises/stack"
	"fmt"
	"math"
	"strconv"
)

var symbolsWeight = map[rune]int{
	'+': 0,
	'-': 0,
	'*': 1,
	'/': 1,
	'^': 2,
	'(': 3,
	')': 3,
}

type operation func(float64, float64) float64

var symbolsOperation = map[rune]operation{
	'+': func(a, b float64) float64 { return a + b },
	'-': func(a, b float64) float64 { return a - b },
	'*': func(a, b float64) float64 { return a * b },
	'/': func(a, b float64) float64 { return a / b },
	'^': func(a, b float64) float64 { return math.Pow(a, b) },
}

func main() {
	input := "6*((5+(2+3^2)*8)+3)"

	output := []rune{}
	stack := stack.NewStack[rune](len(input))
	for _, char := range input {

		// If the character not is a symbol
		if _, ok := symbolsWeight[char]; !ok {
			output = append(output, char)
			continue
		}

		//If the character is ')', we extract all the elements up to the first parenthesis
		if char == ')' {
			bracketsExtract(stack, &output)
			continue
		}

		e := stack.Top()
		// Add the symbol if it has a greater weight than the top symbol of the stack or if the stack is empty
		if stack.Empty() || symbolsWeight[char] > symbolsWeight[e] {
			stack.Push(char)
		} else {
			/* Add the symbol to the output (postfix expression) while making sure that the top symbol is greater than or equal to the input character and different from '(' because it can't be  add the '(' symbol to the postfix expression */

			for !stack.Empty() && e != '(' && symbolsWeight[char] <= symbolsWeight[e] {
				stack.Pop()
				output = append(output, e)
				e = stack.Top()
			}

			//Push the input character to the stack
			stack.Push(char)
		}
	}

	//Add all remaining symbols to the output (postfix expression)
	for !stack.Empty() {
		output = append(output, stack.Pop())
	}
	//Print the postfix expression
	fmt.Printf("%c\n", output)

	//Compute the result of the postifx expression
	fmt.Println(computePostFixExpression(output))
}

// Extract all the symbols that are between the brackets
func bracketsExtract(stack *stack.Stack[rune], output *[]rune) {
	stackElement := stack.Pop()
	for stackElement != '(' {
		*output = append(*output, stackElement)
		stackElement = stack.Pop()
	}
}

// Compute the postifix expression
func computePostFixExpression(postFix []rune) float64 {
	result := stack.NewStack[float64](len(postFix))

	for _, char := range postFix {
		if _, ok := symbolsOperation[char]; ok {
			num1 := result.Pop()
			num2 := result.Pop()

			result.Push(symbolsOperation[char](num2, num1))
		} else {
			value, _ := strconv.ParseFloat(string(char), 64)
			result.Push(value)
		}
	}

	return result.Top()
}
