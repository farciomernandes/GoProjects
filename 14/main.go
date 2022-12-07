package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

// Qualquer coisa que tenha um método "Desativar" pode usar essa função
func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}
	minhaEmpresa := Empresa{}
	Desativacao(wesley)

	//A empresa não é uma pessoa, mas pode usar esse método. Porquê o STRUCT de EMPRESA implementa o método
	// Desativar(), por isso o GO consegue interpretar que EMPRESA implementa a interface PESSOA
	Desativacao(minhaEmpresa)
}
