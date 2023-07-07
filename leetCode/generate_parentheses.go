package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	var recursion func(s string, left int, right int)
	recursion = func(s string, left int, right int) {
		if len(s) == 2*n {
			result = append(result, s)
			return
		}
		if left < n {
			recursion(s+"(", left+1, right)
		}
		if right < left {
			recursion(s+")", left, right+1)
		}
	}
	recursion("", 0, 0)
	return result

}
func main() {
	result := generateParenthesis(5)
	fmt.Println(result)

}
