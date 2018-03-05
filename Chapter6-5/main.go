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

	fizzBuzzFunc := func(mod int, str string) func(n int) (string, bool) {
		return func(n int) (string, bool) {
			if n%mod == 0 {
				return str, true
			}
			return "", false
		}
	}

	fizzBuzzArray := [...]func(n int) (string, bool){
		fizzBuzzFunc(15, "FizzBuzz"),
		fizzBuzzFunc(5, "Buzz"),
		fizzBuzzFunc(3, "Fizzz"),
	}

	for i := 0; i < len(fizzBuzzArray); i++ {
		if str, ok := fizzBuzzArray[i](number); ok {
			return str
		}
	}

	return strconv.Itoa(number)
}
