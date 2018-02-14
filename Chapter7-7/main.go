package main

import "fmt"

func main() {
	myMap := make(map[int]int)
	myMap[1] = 1
	myMap[2] = 2
	//myMap[3] = 4

	fmt.Println(myMap[3])
	if myMap[3] != 0 {
		fmt.Println(myMap[3])
	}

	// exist?
	if value, exist := myMap[3]; exist {
		fmt.Println(value)
	}

}

