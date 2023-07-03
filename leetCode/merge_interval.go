package main

import (
	"fmt"
	"sort"
)

func sortSlices(intervals [][]int) {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

}
func merge(intervals [][]int) [][]int {
	sortSlices(intervals)
	result := make([][]int, 0, len(intervals))
	result = append(result, intervals[0])
	for _, interval := range intervals {
		if top := result[len(result)-1]; top[1] < interval[0] {
			result = append(result, interval)
		} else if interval[1] > top[1] {
			top[1] = interval[1]
		}
	}
	return result

}

func main() {
	firstInterval := []int{1, 2}
	secondInterval := []int{3, 4}
	intervals := [][]int{firstInterval, secondInterval}
	fmt.Println(merge(intervals))

}
