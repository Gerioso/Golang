package main

import "fmt"

type node struct {
	value int
	next  *node
}

func printNodeValue(n *node) {

	fmt.Println(n.value)
	if n.next != nil {
		printNodeValue(n.next)
	}
}

func main() {

	first := node{value: 4}
	second := node{value: 5}
	third := node{value: 6}

	first.next = &second
	second.next = &third

	printNodeValue(&first)

}
