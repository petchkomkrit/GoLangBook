package main

import "fmt"

func main() {
	fizzBuzz(1, 100)
}

func fizzBuzz(start int, end int) {
	for number := start; number <= end; number++ {
		if verifyFizzBuzz(number, 15) {
			showFizzBuzzMessage(number, "FizzBuzz")
		} else if verifyFizzBuzz(number, 5) {
			showFizzBuzzMessage(number, "Fizz")
		} else if verifyFizzBuzz(number, 3) {
			showFizzBuzzMessage(number, "Buzz")
		} else {
			showFizzBuzzMessage(number, "")
		}
	}
}

func showFizzBuzzMessage(number int, msg string) {
	fmt.Println(number, msg)
}

func verifyFizzBuzz(number int, mod int) (result bool) {
	if number%mod == 0 {
		result = true
	}
	return
}
