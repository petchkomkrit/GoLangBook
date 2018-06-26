package main

import (
	"fmt"
)

func main() {
	defer func() {
		x := recover()
		if x != "" {
			fmt.Printf("recovered : %v", x)
		}
	}()

	arr := [...]int{1, 2, 3}
	for i := 0; i < 4; i++ {
		fmt.Printf("%d\n", arr[i])
	}
}
