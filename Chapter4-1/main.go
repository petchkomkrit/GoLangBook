package main

import (
	"fmt"
)

var num int = 56021

func main() {
	fmt.Println("============================")
	var x string = "Hello, world."
	fmt.Println(x)

	var y string 
	y = "Hello, world."
	fmt.Println(y)
	fmt.Println("============================\n")

	fmt.Println("============================")
	z := "Hello, world."
	fmt.Println(z)
	fmt.Printf("z Type : %T\n",z)
	fmt.Println("============================\n")

	fmt.Println("============================")
	fmt.Println(num)
	fmt.Printf("num Type : %T\n",num)
	fmt.Println("============================\n")

	fmt.Println("============================")
	const greeting string = "Hello"
	fmt.Println(greeting)
	//greeting = "Hi"
	fmt.Println("============================\n")

	fmt.Println("============================")
	var(
		a = 5
		b = 10
		c = 20
		s = "Text"
	)
	v1, v2 := "First", "Second"
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(s)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println("============================\n")

	fmt.Println("============================")
	v1, v2 = v2, v1
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println("============================\n")
}

func f() {
	fmt.Println(num)
}