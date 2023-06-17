package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var prefix string
	for _, s := range strs[0] {
		letter := string(s)
		tmp := prefix + letter
		for i := 1; i < len(strs); i++ {
			if !strings.HasPrefix(strs[i], tmp) {
				return prefix
			}
		}
		prefix = tmp
	}
	return prefix
}

func main() {
	s := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(s)
	fmt.Println(result)

}
