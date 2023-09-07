package main

import "fmt"
func main(){
	ch := make(chan int)
	go func(n int, chInt chan int ){
		result := 1
		for i:=1; i <= n; i++{
			result *= i
		}
		fmt.Println(n, "-", result)
		ch <- result
	}(5,ch)
	fmt.Println(<- ch)
	fmt.Println("End")
}