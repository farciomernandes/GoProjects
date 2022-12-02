package main

const a = "Hello, World!"

// Variaveis de escopo global para todos os package main
var (
	b bool
	c int
	d string
	e float64
)

//Variaveis declaradas sem valor inicial, o GO da um valor default

func main() {
	var a string
	//Variavel se escopo local, vista dentro dessa função
	println(a)

	b = 10 //Acusa de erro, é uma linguagem fortmente tipada, por "b" ser boolean, nao deixa setar um valo inteiro
}
