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
	// mod := [3]int{15, 5, 3}
	// msg := [3]string{"FizzBuzz", "Fizz", "Buzz"}

	// for i := 0; i < len(mod); i++ {
	// 	if number%mod[i] == 0 {
	// 		return msg[i]
	// 	}
	// }

	// return strconv.Itoa(number)

	fizzBuzzFunc := func(n int) (string, bool) {
		if n%15 == 0 {
			return "FizzBuzz", true
		}
		return "", false
	}

	buzzFunc := func(n int) (string, bool) {
		if n%5 == 0 {
			return "Buzz", true
		}
		return "", false
	}

	fizzFunc := func(n int) (string, bool) {
		if n%3 == 0 {
			return "Fizz", true
		}
		return "", false
	}

	fizzBuzzArray := [...]func(n int) (string, bool){
		fizzBuzzFunc,
		buzzFunc,
		fizzFunc,
	}

	for i := 0; i < len(fizzBuzzArray); i++ {
		if str, ok := fizzBuzzArray[i](number); ok {
			return str
		}
	}

	return strconv.Itoa(number)
}
