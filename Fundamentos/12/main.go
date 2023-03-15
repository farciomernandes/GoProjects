package main

import (
	"fmt"
)

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

/*
Dessa forma com parentese antes é possível criar um "método da struct"
para ter acesso as propriedades da struct e poder ser acessado a partir
de qualquer variavel do tipo entre parentese.
Porém as alterações feitas são somente locais.
*/
func (cliente Cliente) Desativar() {
	cliente.Ativo = false
	fmt.Printf("O client %s foi desativado\n", cliente.Nome)
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	wesley.Cidade = "São Paulo"
	wesley.Desativar()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", wesley.Nome, wesley.Idade, wesley.Ativo)
}
