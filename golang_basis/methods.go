package main

import "fmt"

type contact struct {
	email string
	phone string
}

type person struct {
	name string
	age  uint16
	contact
}

func (p person) print() {
	fmt.Println("Name:", p.name, "\nAge:", p.age)
	fmt.Println("Contact:", p.email)
}
func main() {
	paul := person{"Paul", 18, contact{"tom@gmail.com", "+1234567899"}}
	paul.print()
	paul.age = 34
	paul.print()
}
