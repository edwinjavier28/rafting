package main

import "fmt"

func main() {
	a := suma(5, 7)
	fmt.Println(a)
}

func suma(a, b int) int {
	return a + b
}

func esPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}
	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}

func encontrarPrimos(inicio, fin int) []int {
	primos := []int{}
	for i := inicio; i <= fin; i++ {
		if esPrimo(i) {
			primos = append(primos, i)
		}
	}
	return primos
}

func main2() {
	inicio := 1
	fin := 100
	primos := encontrarPrimos(inicio, fin)
	fmt.Println("Números primos entre", inicio, "y", fin, ":")
	fmt.Println(primos)
}
