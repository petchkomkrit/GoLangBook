package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(16)
	go incremental(1)
	go incremental(2)
	go incremental(3)
	go incremental(4)
	go incremental(5)
	go incremental(6)
	go incremental(7)
	go incremental(8)
	go incremental(9)
	go incremental(10)
	go incremental(11)
	go incremental(12)
	go incremental(13)
	go incremental(14)
	go incremental(15)
	go incremental(16)

	wg.Wait()
	fmt.Println("Final counter : ", counter)
}

func incremental(n int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		value := counter
		// runtime.Gosched()
		value++
		counter = value
	}
}
