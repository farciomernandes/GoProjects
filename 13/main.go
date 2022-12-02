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
Qualquer Client que possua um método "Desativar" ele automaticamente possui um
implements Pessoa" atachado.
*/
type Pessoa interface {
	Desativar()
}

//Interface só seta metodos e nao atributos

type Empresa struct {
	Nome string
}

func (empresa Empresa) Desativar() {

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

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	wesley.Cidade = "São Paulo"
	minhaEmpresa := Empresa{}
	Desativacao(wesley)
	Desativacao(minhaEmpresa)

}
