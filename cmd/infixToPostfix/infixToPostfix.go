package infixToPostfix

import (
	stack "exercises/packages/stack"
	"math"
	"strconv"
)

var symbolsWeight = map[string]int{
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
	"^": 2,
	"(": 3,
	")": 3,
}

type operation func(float64, float64) float64

var symbolsOperation = map[string]operation{
	"+": func(a, b float64) float64 { return a + b },
	"-": func(a, b float64) float64 { return b - a },
	"*": func(a, b float64) float64 { return a * b },
	"/": func(a, b float64) float64 { return b / a },
	"^": func(a, b float64) float64 { return math.Pow(b, a) },
}

// Infix to postfix
func InfixToPostfix(input []string) []string {
	//Store the postfix expression
	output := []string{}

	//Store the operators and parenthesis
	stack := stack.NewStack[string](len(input))

	for _, char := range input {

		// If the character not is a symbol
		if _, ok := symbolsWeight[char]; !ok {
			output = append(output, char)
			continue
		}

		//If the character is ')', we extract all the elements up to the first parenthesis
		if char == ")" {
			bracketsExtract(stack, &output)
			continue
		}

		e := stack.Top()
		// Add the symbol if it has a greater weight than the top symbol of the stack or if the stack is empty
		if stack.Empty() || (symbolsWeight[char] == symbolsWeight[e] && e == "^") || symbolsWeight[char] > symbolsWeight[e] {
			stack.Push(char)
		} else {
			/* Add the symbol to the output (postfix expression) while making sure that the top symbol is greater than or equal to the input character and different from '(' because it can't be  add the '(' symbol to the postfix expression */

			for !stack.Empty() && e != "(" && symbolsWeight[char] <= symbolsWeight[e] {
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
	return output
}

// Extract all the symbols that are between the brackets
func bracketsExtract(stack *stack.Stack[string], output *[]string) {
	stackElement := stack.Pop()
	for stackElement != "(" {
		*output = append(*output, stackElement)
		stackElement = stack.Pop()
	}
}

// Compute the postifix expression
func ComputePostFixExpression(postFix []string) float64 {
	result := stack.NewStack[float64](len(postFix))

	for _, char := range postFix {
		if _, ok := symbolsOperation[char]; ok {
			num1 := result.Pop()
			num2 := result.Pop()

			result.Push(symbolsOperation[char](num1, num2))
		} else {
			value, _ := strconv.ParseFloat(string(char), 64)
			result.Push(value)
		}
	}

	return result.Top()
}
