package main

import "fmt"

func add(x, y int) (z int) {
	z = x + y
	return
}

func mul(x, y int) (z int) {
	z = x * y
	return
}

func action(v1, v2 int, operator func(int, int) int) {
	var result = operator(v1, v2)
	fmt.Println(result)
}

func main() {
	action(1, 2, add)
	action(1, 2, mul)
}
