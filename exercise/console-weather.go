package main

import (
	"fmt"
	"strconv"
)

func main() {
	weather(91, "Chaing-mai cold")
}

func weather(celsius int, message string) {
	nine := map[string]string{
		"one":   " _ ",
		"two":   "|_|",
		"three": " _|",
	}
	var number [10]map[string]string
	number[9] = nine

	celsiusStr := strconv.Itoa(celsius)
	celsiusStr1 := celsiusStr[:1]
	//celsiusStr2 := celsiusStr[1:]
	celsius1, _ := strconv.Atoi(celsiusStr1)
	//celsius2, _ := strconv.Atoi(celsiusStr2)

	fmt.Println(number[celsius1]["one"])
	fmt.Println(number[celsius1]["two"])
	fmt.Println(number[celsius1]["three"])

	fmt.Println(message)
}
