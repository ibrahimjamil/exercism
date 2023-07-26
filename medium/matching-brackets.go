package medium

import "fmt"

type Stack struct {
	data []rune
}

func (s *Stack) Push(r rune) {
	s.data = append(s.data, r)
}

func (s *Stack) Pop() rune {
	if len(s.data) == 0 {
		return 0
	}

	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return top
}

func isValidCombination(combination string) bool {
	stack := &Stack{data: make([]rune, 0)}

	for _, char := range combination {
		switch char {
		case '(':
			stack.Push(char)
		case '{':
			stack.Push(char)
		case '[':
			stack.Push(char)
		case ')':
			if stack.Pop() != '(' {
				return false
			}
		case '}':
			if stack.Pop() != '{' {
				return false
			}
		case ']':
			if stack.Pop() != '[' {
				return false
			}
		}
	}

	return len(stack.data) == 0
}

func MatchingBrackets(combination string) {
	if isValidCombination(combination) {
		fmt.Printf("correctly placed")
	} else {
		fmt.Printf("not correctly placed")
	}
}
