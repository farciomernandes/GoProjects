package main

import "fmt"

func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

/*
QUANDO NÃO USAR PONTEIRO: {
  - Quando só precisar de uma cópia dos dados para executar um método
  - Quando não deseja alterar o valor em memória durante alguma operação
  - Quando não quer que os valores passados para o método sejam mutaveis
    }

QUANDO USAR PONTEIRO: {
  - Quando você quer que o método possa manipular de forma global o valor das informações enviadas
    -
    }
*/
func main() {
	valor1 := 10
	valor2 := 20
	soma(&valor1, &valor2)
	fmt.Println(valor1)
	fmt.Println(valor2)
}
