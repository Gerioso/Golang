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

func main() {
	paul := person{"Paul", 18, contact{"tom@gmail.com", "+1234567899"}}
	paul.age = 34
	fmt.Println(paul.name, paul.age)
	fmt.Println(paul.email)
}
