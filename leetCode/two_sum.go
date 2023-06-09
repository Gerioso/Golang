package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		m[v] = k
	}

	for k, v := range nums {
		diffs := target - v
		position, ok := m[diffs]
		if ok && k != position {
			return []int{k, position}
		}
	}
	return []int{0}

}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)

}
