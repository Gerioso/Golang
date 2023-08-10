package main

import (
	"fmt"
	"strconv"
)

func intToDigit(k int) []byte {
	strNumber := strconv.Itoa(k)
	digitK := make([]byte, len(strNumber))
	for digit := range strNumber {
		digitK[digit] = strNumber[digit]
	}
	return digitK
}
func compress(chars []byte) int {
	i := 0
	for i < len(chars)-1 {
		if chars[i] == chars[i+1] {
			k := 1
			for i+k < len(chars) && chars[i] == chars[i+k] {
				k++
			}
			digitK := intToDigit(k)
			chars = append(chars[:i+1], append(digitK, chars[i+k:]...)...)
			i += len(digitK) + 1
		} else {
			i++
		}
	}
	return len(chars)
}
func main() {
	result := compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'})
	fmt.Println(result)

}
