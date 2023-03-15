package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("itens=%d capacidade=%d %v\n", len(s), cap(s), s)

	// s[:0] = Todo valor no array, só mostre os valores do inicio ate a posicao "0" que eu digitei, o resto apague ( a capacidade existe, mas o array vai mostrar vazio )
	fmt.Printf("itens=%d capacidade=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	// s[:4] = Mostra os 4 primeiros valores do SLICE e o resto ele ignora( a capacidade existe, mas o array vai mostrar vazio )
	fmt.Printf("itens=%d capacidade=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	// s[2:] = Mostra todos os valores do SLICE, ignorando do 0 até a posicao "2"( a capacidade existe, mas o array vai mostrar vazio )
	fmt.Printf("itens=%d capacidade=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	//append = Dobra o tamanho inicial, mesmo que nao use tudo
	s = append(s, 110)
	fmt.Printf("itens=%d capacidade=%d %v\n", len(s), cap(s), s)

}
