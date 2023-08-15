package main

import "fmt"

func restoreString(s string, indices []int) string {
	result := make([]byte, len(indices))
	for shuff, i := range indices {
		result[i] = s[shuff]
	}
	return string(result)
}

func main() {
	result := restoreString("asd", []int{0, 2, 1})
	fmt.Println(string(result))

}
