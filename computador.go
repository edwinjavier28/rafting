package main

import (
	"fmt"
	"sort"
)

func nueve() {}
func dies()  {}
func once()  {}

func main3() {
	numeros := []int{4, 2, 7, 1, 9, 5, 3, 8, 6}
	fmt.Println("Array sin ordenar:", numeros)

	sort.Ints(numeros) // Ordenar el array de enteros

	fmt.Println("Array ordenado:", numeros)
}
