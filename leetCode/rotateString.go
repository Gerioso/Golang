package main

import "fmt"

func rotateString(s string, goal string) bool {
	check := false
	for i := 0; i < len(goal); i++ {
		if s == goal {
			check = true
			break
		}
		s = s[1:] + string(s[0])
	}
	return check
}

func main() {
	result := rotateString("asd", "asd")
	fmt.Println(result)

}
