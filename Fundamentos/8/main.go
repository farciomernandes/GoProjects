package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(2, 3, 5, 40))
}

// Enviando "N" valores para uma função
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
