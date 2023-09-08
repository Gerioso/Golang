package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var (
		formatString = strings.ToLower(strings.ReplaceAll(s, " ", ""))
		list         = []rune(formatString)
		first        = 0
		last         = len(list) - 1
	)
	for first < last {
		left, right := list[first], list[last]
		if left == right {
			first++
			last--
		} else if !unicode.IsLetter(left) && !unicode.IsNumber(left) {
			first++
		} else if !unicode.IsLetter(right) && !unicode.IsNumber(right) {
			last--
		} else {
			return false
		}
	}
	return true

}

func main() {
	result := isPalindrome("0P")
	fmt.Println(result)

}
