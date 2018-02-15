package main

import "fmt"

func main() {
	for position, ascii := range "golang" {
		fmt.Println(position, ascii)
		fmt.Printf("%v\n", string(ascii))
	}
}
