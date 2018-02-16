package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(weatherCelcius(25, "Bangkok few cloud"))
	fmt.Println(weatherCelcius(34, "Tak sunny"))
	fmt.Println(weatherCelcius(17, "Phuket rainy"))
	fmt.Println(weatherCelcius(9, "Chaing-mai cold"))
}

func weatherCelcius(celsius int, message string) string {

	var number [11]map[string]string
	number[0] = map[string]string{
		"firstLine":  " _ ",
		"secondLine": "| |",
		"thirdLine":  "|_|",
	}
	number[1] = map[string]string{
		"firstLine":  "   ",
		"secondLine": "  |",
		"thirdLine":  "  |",
	}
	number[2] = map[string]string{
		"firstLine":  " _ ",
		"secondLine": " _|",
		"thirdLine":  "|_ ",
	}
	number[3] = map[string]string{
		"firstLine":  " _ ",
		"secondLine": " _|",
		"thirdLine":  " _|",
	}
	number[4] = map[string]string{
		"firstLine":  "   ",
		"secondLine": "|_|",
		"thirdLine":  "  |",
	}
	number[5] = map[string]string{
		"firstLine":  " _ ",
		"secondLine": "|_ ",
		"thirdLine":  " _|",
	}
	number[6] = map[string]string{
		"firstLine":  " _ ",
		"secondLine": "|_ ",
		"thirdLine":  "|_|",
	}
	number[7] = map[string]string{
		"firstLine":  " _ ",
		"secondLine": "  |",
		"thirdLine":  "  |",
	}
	number[8] = map[string]string{
		"firstLine":  " _ ",
		"secondLine": "|_|",
		"thirdLine":  "|_|",
	}
	number[9] = map[string]string{
		"firstLine":  " _ ",
		"secondLine": "|_|",
		"thirdLine":  " _|",
	}
	number[10] = map[string]string{
		"firstLine":  "   ",
		"secondLine": " _ ",
		"thirdLine":  "   ",
	}

	celsiusStr := strconv.Itoa(celsius)

	var line [3]string
	for i := 1; i <= len(celsiusStr); i++ {
		celsiusExtractStr := celsiusStr[i-1 : i]

		var celsiusNumber int
		if celsiusExtractStr == "-" {
			celsiusNumber = 10
		} else {
			celsiusNumber, _ = strconv.Atoi(celsiusExtractStr)
		}

		line[0] = line[0] + number[celsiusNumber]["firstLine"]
		line[1] = line[1] + number[celsiusNumber]["secondLine"]
		line[2] = line[2] + number[celsiusNumber]["thirdLine"]
	}
	line[2] = line[2] + " c"

	for i, _ := range line {
		fmt.Println(line[i])
	}

	return message
}
