package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	var i, j, k int = m - 1, n - 1, m + n - 1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	return nums1
}

func main() {
	nums1 := []int{1, 3, 5, 0, 0}
	nums2 := []int{2, 4}
	result := merge(nums1, 3, nums2, 2)
	fmt.Println(result)

}
