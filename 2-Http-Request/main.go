package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("http://google.com/")

	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	println(string(res))
	// Quando abre um arquivo é preciso fechar
	req.Body.Close()
}
