package main

import (
	"fmt"
)

func moveZeroes(nums []int) []int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && i != index {
			nums[index] = nums[i]
			nums[i] = 0
			index++
		}
	}
	return nums
}

func main() {
	result := moveZeroes([]int{1, 0})
	fmt.Println(result)

}
