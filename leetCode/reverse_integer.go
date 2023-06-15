package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	maxPositive := math.MaxInt32
	maxNegative := math.MinInt32
	sign := 1
	result := 0
	if x == 0 {
		return 0
	}
	if x < 0 {
		sign = -1
		x = -x
	}
	for x > 0 {
		div := x % 10
		result *= 10
		result += div
		x /= 10
	}
	result *= sign
	if result > maxPositive || result < maxNegative {
		return 0
	}
	return result

}

func main() {
	result := reverse(-123)
	fmt.Println(result)

}
