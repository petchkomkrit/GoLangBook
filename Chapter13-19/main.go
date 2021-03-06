package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	print(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(int <-chan int, out chan<- int) {
	for x := range int {
		out <- x * x
	}
	close(out)
}

func print(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
