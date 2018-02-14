package main

import "fmt"

func main() {
	slice := make([]int, 3, 10)
	slice[0] = 5
	slice[1] = 10
	slice[2] = 15
	fmt.Println(slice)

	slice2 := []int{1,2,3,4,5}
	fmt.Println(slice2)

	fmt.Println("\nSlice with length and capacity")
	fmt.Printf("slice : lenght %v, capacity %v, %v\n",len(slice), cap(slice), slice)

	// append
	for i := 4; i< 12; i++ {
		slice = append(slice, i)
	}
	fmt.Printf("slice : lenght %v, capacity %v, %v\n",len(slice), cap(slice), slice)

	for i := 6; i< 12; i++ {
		slice2 = append(slice2, i)
	}
	fmt.Printf("slice2 : lenght %v, capacity %v, %v\n",len(slice2), cap(slice2), slice2)

}

