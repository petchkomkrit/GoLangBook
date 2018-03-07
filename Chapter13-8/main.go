package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	var a, b value
	var wg sync.WaitGroup
	wg.Add(2)
	go printSum(&a, &b, &wg)
	go printSum(&b, &a, &wg)
	wg.Wait()
}

func printSum(objectA, objectB *value, wg *sync.WaitGroup) {
	defer wg.Done()
	objectA.mu.Lock()
	defer objectA.mu.Unlock()

	time.Sleep(2 * time.Second)
	objectB.mu.Lock()
	defer objectB.mu.Unlock()

	fmt.Printf("sum=%v\n", objectA.value+objectB.value)
}
