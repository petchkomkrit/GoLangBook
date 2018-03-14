package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	volumn := 200
	start := time.Now()

	container := order(volumn)
	for _, cup := range container {
		fmt.Println(cup)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func order(volumn int) (container []string) {

	order := make(chan string)
	order1 := make(chan string)
	order2 := make(chan string)

	go cashier(order, volumn)
	go barista(order, order1, volumn)
	go waiter(order1, order2)

	for orderFinish := range order2 {
		//orderFinish := <-order2
		container = append(container, orderFinish)
	}

	return
}

func cashier(orderOut chan<- string, volumn int) {
	// cashier
	for i := 1; i <= volumn; i++ {
		time.Sleep(5 * time.Millisecond)
		coffee := fmt.Sprintf("order : %d", i)
		orderOut <- coffee
	}
	close(orderOut)
}

func barista(orderIn <-chan string, orderOut chan<- string, volumn int) {

	var wgBarista sync.WaitGroup
	wgBarista.Add(4)

	wg.Add(volumn)
	for coffee := range orderIn {
		//barista
		go brew(coffee, orderOut)
	}
	wg.Wait()
	close(orderOut)
}

func waiter(orderIn <-chan string, orderOut chan<- string) {

	for coffee := range orderIn {
		//waiter
		time.Sleep(5 * time.Millisecond)
		coffee = fmt.Sprintf("%s %s", coffee, "ready :)")
		orderOut <- coffee
	}
	close(orderOut)
}

func brew(orderIn string, orderOut chan<- string) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	coffee := fmt.Sprintf("%s %s", orderIn, "Espresso")
	orderOut <- coffee
}
