package main

import "fmt"

type Semaphore chan struct{}

func (s Semaphore) Inc(n int) {
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}
}

func (s Semaphore) Decr(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

func main() {
	const number = 5
	s := make(Semaphore, number)
	for i := 0; i < number; i++ {
		go func(num int) {
			fmt.Println(num)
			defer s.Inc(1)
		}(i)
	}
	s.Decr(number)

}
