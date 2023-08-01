package main

import (
	"fmt"
)

type Key [26]byte

func generateKey(str string) (key Key) {
	for i := range str {
		key[str[i]-'a']++
	}
	return
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[Key][]string)
	for _, str := range strs {
		key := generateKey(str)
		groups[key] = append(groups[key], str)
	}
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result

}

func main() {
	s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(s)
	fmt.Println(result)

}
