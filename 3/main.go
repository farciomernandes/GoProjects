package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Primeira linha")
	fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")

	//"defer" atrasa a execução do "req.Body.Close()"
	//"defer" faz a execução do código pular essa linha, e só utilizar após todo o restante do código executar

}
