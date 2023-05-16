package main

import "fmt"

func main() {

	var x int = 4
	var p *int = &x            // указатель получает адрес переменной
	fmt.Println("Address:", p) // значение указателя - адрес переменной x
	fmt.Println("Value:", *p)  // значение переменной x
	f := 2.3
	pf := &f

	fmt.Println("Address:", pf)
	fmt.Println("Value:", *pf)
}
