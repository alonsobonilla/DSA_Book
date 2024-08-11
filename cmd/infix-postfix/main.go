package main

import (
	stack "exercises/stack"
)

var signs = map[rune]int{
	'+': 0,
	'-': 0,
	'*': 1,
	'/': 1,
	'^': 2,
	'(': 10,
	')': 11,
}

func main() {
	input := "a+b*c+(2^(2^3))*g"
	signsStack := stack.NewStack[rune](len(input))
	output := make([]rune, len(input))

	for _, element := range input {

		if valuePrecedence, ok := signs[element]; ok {

			top, ok := signsStack.Top()
			if ok {
				if signs[top] > valuePrecedence {
					signsStack.Push(element)
				} else {
					
				}
			} else {
				signsStack.Push(element)
			}
		} else {
			output = append(output, element)
		}

	}
}
