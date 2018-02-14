package main

import "fmt"

func main() {
	fmt.Println(add(1,2,3,4))	

	xs := []int{1,2,3,5}
	fmt.Println(add(xs...))	
}

func add(args ...int) (int, int) {
	total := 0
	total1 := 0

	for _,value := range args { // range founction return index, value of slice
		total += value
	}

	for i:=0;i<len(args);i++ {
		total1 += args[i]
	}
	return total, total1
}
