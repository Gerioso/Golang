package main

import "fmt"

func romanToInt(s string) int {
	var sum, cur, prev int
	m := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := len(s) - 1; i >= 0; i-- {
		cur = m[s[i]]
		if cur < prev {
			sum -= cur
		} else {
			sum += cur
		}
		prev = cur

	}
	return sum

}

func main() {
	s := "MCMXCIV"
	result := romanToInt(s)
	fmt.Println(result)

}
