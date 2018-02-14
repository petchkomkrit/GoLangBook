package main

import (
	"fmt"
	"strconv"
)

func main() {
	for number := 1; number <= 100; number++ {
		fmt.Println(fizzBuzz(number))
	}
}

func fizzBuzz(number int) string {
	var msg string
	if verifyFizzBuzz(number, 15) {
		msg = "FizzBuzz"
	} else if verifyFizzBuzz(number, 5) {
		msg = "Fizz"
	} else if verifyFizzBuzz(number, 3) {
		msg = "Buzz"
	} else {
		msg = strconv.Itoa(number)
	}

	return msg
}

func verifyFizzBuzz(number int, mod int) (result bool) {
	if number%mod == 0 {
		result = true
	}
	return
}
