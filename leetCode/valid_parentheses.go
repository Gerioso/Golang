package main

import (
	"fmt"
)

type Stack []rune

func (s *Stack) Push(val rune) {
	*s = append(*s, val)
}
func (s *Stack) Pop() (bool, rune) {
	if len(*s) == 0 {
		return false, ' '
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return true, value

}

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	rules := map[rune]rune{
		'[': ']',
		'{': '}',
		'(': ')',
	}
	var stack Stack

	for _, v := range s {
		if _, ok := rules[v]; ok {
			stack.Push(v)
			fmt.Println(1)
		} else if check, value := stack.Pop(); !check || v != rules[value] {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	s := "(("
	result := isValid(s)
	fmt.Println(result)

}
