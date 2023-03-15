package main

import (
	"fmt"
)

func main() {

	total := func() int {
		return sum(2, 3, 5, 40) * 2
	}()

	fmt.Println(total)
}

// Enviando "N" valores para uma função
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
