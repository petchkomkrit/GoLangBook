package main

import (
	"fmt"
	"time"
)

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
	go barista(order, order1)
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

func barista(orderIn <-chan string, orderOut chan<- string) {

	for coffee := range orderIn {
		//barista
		//coffee := <-orderIn
		time.Sleep(100 * time.Millisecond)
		coffee = fmt.Sprintf("%s %s", coffee, "Espresso")
		orderOut <- coffee
	}
	close(orderOut)
}

func waiter(orderIn <-chan string, orderOut chan<- string) {

	for coffee := range orderIn {
		//waiter
		//coffee := <-orderIn
		time.Sleep(5 * time.Millisecond)
		coffee = fmt.Sprintf("%s %s", coffee, "ready :)")
		orderOut <- coffee
	}
	close(orderOut)
}
