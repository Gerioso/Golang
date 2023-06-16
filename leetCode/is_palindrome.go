package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var reverse int
	init := x
	for x > 0 {
		reverse = reverse*10 + x%10
		x /= 10
	}
	return init == reverse

}

func main() {
	result := isPalindrome(121)
	fmt.Println(result)

}
