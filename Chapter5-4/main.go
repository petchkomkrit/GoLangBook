package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10)
	//randomNumber := 5
	
	for count := 0; count <= 4; count++ {
		var number int
		
		fmt.Print("\nEnter your number : ")
		fmt.Scanf("%v\n", &number)
		fmt.Print("Your number is : ",number," -> ")

		if number == randomNumber {
			fmt.Println("Correct !!")
			fmt.Println("Random number is :",randomNumber)
			return
		} else if number > randomNumber {
			fmt.Println("มากกว่า")
		} else if number < randomNumber {
			fmt.Println("น้อยกว่า")
		}
	}

	
}
