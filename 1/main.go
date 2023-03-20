package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	//Salva string
	/*tamanho, err := f.WriteString("Hello World!")

	if err != nil {
		panic(err)
	} */

	//Salva dados
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))

	fmt.Println("Arquivo criado com sucesso! Tamanho %d bytes.", tamanho)

	f.Close()

	//Leitura de arquivos (Lê como bytes)
	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	//Imprime os valores e converte de bytes para string
	fmt.Println(string(arquivo))

	//Leitura pouco a pouco do arquivo
	// ----------------------------------------------------------------
	arquivo2, err := os.Open("arquivo.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2) // cria o reader capaz de ler o arquivo
	buffer := make([]byte, 10)          // buffer de 10 bytes

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		// Imprime e junta os pedaços em um slice
		fmt.Println(buffer[:n])
	}
	// ----------------------------------------------------------------

	//Apaga um arquivo
	err := os.Remove("arquivo.txt")

	if err != nil {
		panic(err)
	}

}
