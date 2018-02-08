package main

import (
	"fmt"
)

func main() {
	fmt.Print("Exnter Number : ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println("Multiple 2 of input is :", output)
}
