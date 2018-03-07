package main

import (
	"fmt"
)

func main() {
	// c := make(chan int)
	// go func() { c <- 1 }()
	// go func() { c <- 2 }()
	// go func() { c <- 3 }()
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// Fix by use buffered channel
	c := make(chan int, 2)
	go func() { c <- 1 }()
	go func() { c <- 2 }()
	go func() { c <- 3 }()
	fmt.Println(<-c)
	fmt.Println(<-c)
}
