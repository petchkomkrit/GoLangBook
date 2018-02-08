package main

import (
	"fmt"
	"math"
)


func main() {
	fmt.Println("===== Floating point =====")
	third := 1.0/3.0
	fmt.Printf("Third = %v\n", third)
	fmt.Printf("Third + Third + Third = %v\n", third+third+third)

	fmt.Println("===== Comparing floating point =====")
	fmt.Println("0.1 + 0.2 == 0.3 is ", 0.1+0.2 == 0.3)
	var num float64
	//var num64 float64
	num = 0.1
	num += 0.2
	//num64 = 0.1
	//num64 += 0.2
	fmt.Println("num == 0.3 is ", num == 0.3)
	fmt.Println("num is ", num)
	//fmt.Println("num is ", num64)


	fmt.Printf("\n\nfloat32(0.1): %064b\n", math.Float32bits(0.1))
	fmt.Printf("float32(0.2): %064b\n", math.Float32bits(0.2))
	fmt.Printf("float64(0.1): %064b\n", math.Float64bits(0.1))
	fmt.Printf("float64(0.2): %064b\n", math.Float64bits(0.2))
}