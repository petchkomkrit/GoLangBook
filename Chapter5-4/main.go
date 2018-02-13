package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(5)
	//randomNumber := 5
	
	for count := 0; count <= 4; count++ {
		var number int
		
		fmt.Print("Enter your number : ")
		fmt.Scanf("%v\n", &number)
		fmt.Println("Your number is :",number)

		if number == randomNumber {
			fmt.Println("Correct !!")
			return
		} else if number > randomNumber {
			fmt.Println("มากกว่า")
		} else if number < randomNumber {
			fmt.Println("น้อยกว่า")
		}
	}

	fmt.Println("Random number is :",randomNumber)
}
