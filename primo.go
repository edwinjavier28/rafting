package main

import (
	"fmt"
)

func esPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}
	if numero <= 3 {
		return true
	}
	if numero%2 == 0 || numero%3 == 0 {
		return false
	}
	i := 5
	for i*i <= numero {
		if numero%i == 0 || numero%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func primos() {
	numeros := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for _, numero := range numeros {
		if esPrimo(numero) {
			fmt.Printf("%d es primo\n", numero)
		} else {
			fmt.Printf("%d no es primo\n", numero)
		}
	}
}
