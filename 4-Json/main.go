package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json: numero` //Singifica que um "numero" no JSON se refere ao Saldo
	Saldo  int `json: saldo`  //Singifica que um "saldo" no JSON se refere ao Saldo
	//Pode colocar `json: -` parar ignorar algum atributo especifico
}

func main() {

	// CONVERTENDO STRUCT EM JSON

	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta) //Transforma STRUCT em JSON, mas retorna em BYTE

	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)

	if err != nil {
		println(err)
	}
	// ou -> encoder := json.NewEncoder(os.Stdout) + encoder.Encode(conta)

	// -----------------------------------------------------------------------

	// CONVERTENDO JSON EM STRUCT

	jsonPuro := []byte(`"Numero":2, "Saldo": 150`)
	var contaX Conta
	//Converter JSON em STRUCT
	err = json.Unmarshal(jsonPuro, &contaX) // Envia o endereço na memória no lugar do nome da variável

	if err != nil {
		println(err)
	}

	println(contaX.Saldo)
}
