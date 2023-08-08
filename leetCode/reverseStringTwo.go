package main

import "fmt"

func reverseString(s string) string {
	lenght := len(s)
	runes := []rune(s)
	for i := 0; i < lenght/2; i++ {
		runes[i], runes[lenght-1-i] = runes[lenght-1-i], runes[i]
	}
	return string(runes)

}

func reverseStr(s string, k int) string {
	ans := ""
	if len(s) < 2 {
		ans = s
		return ans
	}
	if len(s) < k {
		ans = reverseString(s)
		return ans
	}
	i := 0
	for i < len(s) {
		if i+2*k < len(s) {
			ans += reverseString(s[i:i+k]) + s[i+k:i+k*2]
			i += 2 * k
		} else if i+k < len(s) {
			ans += reverseString(s[i:i+k]) + s[i+k:]
			break
		} else {
			ans += reverseString(s[i:])
			break
		}

	}
	return ans
}

func main() {
	result := reverseStr("abcdefg", 2)
	fmt.Println(string(result))

}
