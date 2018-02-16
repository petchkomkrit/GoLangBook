package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(weather(9, "Chaing-mai cold"))
}

func weather(celsius int, message string) string {
	one := map[string]string{
		"one":   "   ",
		"two":   "  |",
		"three": "  |",
	}
	two := map[string]string{
		"one":   " _ ",
		"two":   " _|",
		"three": "|_ ",
	}
	three := map[string]string{
		"one":   " _ ",
		"two":   " _|",
		"three": " _|",
	}
	four := map[string]string{
		"one":   "   ",
		"two":   "|_|",
		"three": "  |",
	}
	five := map[string]string{
		"one":   " _ ",
		"two":   "|_ ",
		"three": " _|",
	}
	six := map[string]string{
		"one":   " _ ",
		"two":   "|_ ",
		"three": "|_|",
	}
	seven := map[string]string{
		"one":   " _ ",
		"two":   "  |",
		"three": "  |",
	}
	eight := map[string]string{
		"one":   " _ ",
		"two":   "|_|",
		"three": "|_|",
	}
	nine := map[string]string{
		"one":   " _ ",
		"two":   "|_|",
		"three": " _|",
	}
	zero := map[string]string{
		"one":   " _ ",
		"two":   "| |",
		"three": "|_|",
	}
	var number [10]map[string]string
	number[1] = one
	number[2] = two
	number[3] = three
	number[4] = four
	number[5] = five
	number[6] = six
	number[7] = seven
	number[8] = eight
	number[9] = nine
	number[0] = zero

	celsiusStr := strconv.Itoa(celsius)
	celsiusLenght := len(celsiusStr)

	var line [3]string
	for i := 1; i <= celsiusLenght; i++ {
		celsiusExtractStr := celsiusStr[i-1 : i]
		celsiusNumber, _ := strconv.Atoi(celsiusExtractStr)

		line[0] = line[0] + number[celsiusNumber]["one"]
		line[1] = line[1] + number[celsiusNumber]["two"]
		line[2] = line[2] + number[celsiusNumber]["three"]
	}
	line[2] = line[2] + " c"

	for i, _ := range line {
		fmt.Println(line[i])
	}

	return message
}
