package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	w float64
	h float64
}

type Rectangle1 struct {
	w float64
	h float64
}

type Circle struct {
	x float64
	y float64
	r float64
}

func (r Rectangle) area() float64 {
	return r.w * r.h
}

func (r Rectangle1) area() float64 {
	return r.w * r.h * 2
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type measure interface {
	area() float64
}

func printArea(m measure) {
	fmt.Println(m.area())
}

func main() {
	c := Circle{0, 0, 5}
	printArea(c)

	r := Rectangle{3, 4}
	printArea(r)

	r1 := Rectangle1{3, 4}
	printArea(r1)
}
