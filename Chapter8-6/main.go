package main

import "fmt"

func main() {
	maps := map[int]int{1: 1, 2: 2, 3: 3}
	double(maps)
	fmt.Printf("original address %p\n", maps)
	fmt.Printf("original %v\n", maps)
}

func double(nums map[int]int) {
	fmt.Printf("double addr %p\n", nums)
	for i := range nums {
		nums[i] *= 2
	}
	fmt.Println(nums)
}
