package main

import "fmt"

func main() {
	a := suma(5, 7)
	fmt.Println(a)
}

func suma(a, b int) int {
	return a + b
}
