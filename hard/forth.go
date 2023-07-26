package hard

import (
	"strconv"
	"strings"
)

func Forth(input string) int {
	stack := []int{}
	words := strings.Fields(input)
	definedWords := make(map[string]string)

	for i := 0; i < len(words); i++ {
		word := words[i]
		switch strings.ToUpper(word) {
		case "+":
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case "-":
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a-b)
		case "*":
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "/":
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a/b)
		case "DROP":
			stack = stack[:len(stack)-1]
		case "SWAP":
			stack[len(stack)-1], stack[len(stack)-2] = stack[len(stack)-2], stack[len(stack)-1]
		case ":":
			i += 2
			newWord := words[i+1]
			definition := ""

			for ; words[i] != ";" && i < len(words); i++ {
				definition += words[i] + " "
			}

			definedWords[newWord] = definition
		default:
			if val, err := strconv.Atoi(word); err == nil {
				stack = append(stack, val)
			}
		}
	}

	return stack[len(stack)-1]
}
