package main

import "fmt"

func main() {
	fmt.Println("===== String =====")
	backticks := `hello world!,
today's goos day`
	fmt.Println(backticks)
	
	doubleQuotes := "hello world!,\ntoday's goos day."
	fmt.Println(doubleQuotes)
}