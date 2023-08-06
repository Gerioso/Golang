package main

import "fmt"

func reverseString(s []byte) []byte {
	lenght := len(s)
	for i := 0; i < lenght/2; i++ {
		s[i], s[lenght-1-i] = s[lenght-1-i], s[i]
	}
	return s

}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	result := reverseString(s)
	fmt.Println(string(result))

}
