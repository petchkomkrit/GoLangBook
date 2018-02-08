package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter Fahrenheit : ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)
	//celsius := (fahrenheit-32)*(5/9)
	celsius := (fahrenheit-32)*5/9
	fmt.Printf("Celsius is : %.2f", celsius)
}
