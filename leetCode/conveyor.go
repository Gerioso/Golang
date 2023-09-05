package main

import "fmt"

func main() {
	data := make(chan int)
	doubleData := make(chan int)

	go func() {
		for i := 0; i < 9; i++ {
			data <- i
		}
		close(data)
	}()
	go func() {
		for val := range data {
			doubleData <- val * 2
		}
		close(doubleData)
	}()

	for i := range doubleData {
		fmt.Println(i)
	}
}
