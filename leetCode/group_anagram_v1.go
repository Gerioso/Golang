package main

import (
	"fmt"
	"sort"
	"strings"
)

func compareLetters(str1 string, str2 string) bool {
	slice1 := strings.Split(str1, "")
	slice2 := strings.Split(str2, "")
	sort.Strings(slice1)
	sort.Strings(slice2)
	sortedStr1 := strings.Join(slice1, "")
	sortedStr2 := strings.Join(slice2, "")
	return sortedStr1 == sortedStr2
}

func checkAnagram(str string, groups [][]string) bool {

	for i := 0; i < len(groups); i++ {

		if len(str) == len(groups[i][0]) && compareLetters(str, groups[i][0]) {
			groups[i] = append(groups[i], str)
			return true
		}
	}
	return false

}

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0, len(strs))
	for _, str := range strs {
		if !checkAnagram(str, result) {
			result = append(result, []string{str})
		}
	}
	return result

}

func main() {
	s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(s)
	fmt.Println(result)

}
