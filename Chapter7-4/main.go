package main

import "fmt"

func main() {
	slice := []int{10, 20, 30}
	fmt.Println(slice)

	newSlice := make([]int ,2)
	fmt.Println(newSlice)

	copy(slice, newSlice)
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("slice: %v\n", newSlice)
}

