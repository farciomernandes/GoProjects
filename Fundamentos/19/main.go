package main

import "fmt"

func main() {

	var minhaVar interface{} = "Wesley Willians"

	//Pega o valor de uma interface vazia e seta como string
	println(minhaVar.(string))

	//Tenta converter o valor de uma interface, caso dê sucesso o segundo parametro vem true e o primeiro vem o resultado
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)

	// Vai retornar um PANIC pq não está conseguindo converter
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v", res2)

}
