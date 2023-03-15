package main

import (
	"fmt"

	"github.com/farciomernandes/goexpert/matematica"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("Resultado: ", s)
	fmt.Println("UUID criado: ", uuid.New())
}
