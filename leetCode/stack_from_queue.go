package main

import (
	"fmt"
)

type Queue struct {
	items []int
}

func newQueue() *Queue {
	return &Queue{items: make([]int, 0)}
}

func (q *Queue) push(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) pop() int {
	top := q.items[0]
	q.items = q.items[1:]
	return top
}

func (q Queue) Empty() bool {
	return len(q.items) == 0
}

type MyStack struct {
	queue *Queue
}

func Constructor() MyStack {
	return MyStack{queue: newQueue()}
}

func (this *MyStack) Push(x int) {
	reverso := newQueue()
	reverso.push(x)
	for !this.queue.Empty() {
		y := this.queue.pop()
		reverso.push(y)

	}
	this.queue = reverso

}

func (this *MyStack) Pop() int {
	return this.queue.pop()
}

func (this *MyStack) Top() int {
	return this.queue.items[0]
}

func (this *MyStack) Empty() bool {
	return this.queue.Empty()

}

func main() {
	myStack := Constructor()
	myStack.Push(1)
	myStack.Push(2)
	fmt.Println(myStack.queue.items)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Empty())

}
