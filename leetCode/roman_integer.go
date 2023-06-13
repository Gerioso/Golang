package main

import "fmt"

func romanToInt(s string) int {
	m := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0

	for i := 0; i < len(s); i++ {
		if i+1 != len(s) && m[s[i]] < m[s[i+1]] {
			sum -= m[s[i]]
		} else {
			sum += m[s[i]]
		}

	}
	return sum

}

func main() {
	s := "MCMXCIV"
	result := romanToInt(s)
	fmt.Println(result)

}
