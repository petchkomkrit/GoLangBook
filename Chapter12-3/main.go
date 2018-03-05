package main

import (
	"fmt"
)

func main() {
	addFunc := func(a int) func(b int) int {
		return func(b int) int {
			return a + b
		}
	}

	addTwoWith := addFunc(2)
	fmt.Println(addTwoWith(3))
}
