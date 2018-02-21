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
	mod := [3]int{15, 5, 3}
	msg := [3]string{"FizzBuzz", "Fizz", "Buzz"}

	for i := 0; i < len(mod); i++ {
		if number%mod[i] == 0 {
			return msg[i]
		}
	}

	return strconv.Itoa(number)
}
