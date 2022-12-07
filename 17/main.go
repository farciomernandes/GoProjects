package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo int
}

func (c Cliente) andou() {
	c.nome = "Wesley Willians"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

/*
ao falar (c *Conta) fala que é para mudar o valor no escopo global
e não no escopo local
*/
func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {
	wesley := Cliente{
		nome: "Wesley",
	}
	wesley.andou()
	fmt.Printf("O valor da struct com nome é %v\n", wesley.nome)

	conta := Conta{saldo: 100}
	conta.simular(200)
	fmt.Printf("O valor que esta na sua carteira é %d \n", conta.saldo)
}
