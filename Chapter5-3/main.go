package main

import (
	"fmt"
)

func main() {
	for number := 1; number <= 100; number++ {
		switch {
		case number%15 == 0:
			fmt.Println("FizzBuzz")
		case number%3 == 0:
			fmt.Println("Fizz")
		case number%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(number)
		}
	}
}
