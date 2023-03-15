package main

import "fmt"

//import "fmt"

func main() {

	/*
		Variavel aponta -> Ponteiro aponta -> Valor na memoria
	*/

	a := 10

	//Cria um ponteiro do tipo inteiro que aponta para o VALOR NA MEMORIA que a o ponteiro da variavel A aponta
	var ponteiro *int = &a
	fmt.Println(ponteiro)

	//Altera o valor da variavel a qual esse ponteiro está apontando;
	*ponteiro = 20
	fmt.Println(a)

	// Cria um ponteiro chamado B que aponta para o "endereço" que guarda o valor na memoria de A
	b := &a

	// Imprime o endereço
	fmt.Println(b)

	// Imprime o valor dentro do endereço
	fmt.Println(*b)

	// Muda o valor guardado no endereço que B aponta
	*b = 150

	fmt.Println(a)
}
