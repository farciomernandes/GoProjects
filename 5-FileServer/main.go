package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServe := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fileServe)

	log.Fatal(http.ListenAndServe(":8080", mux))

}
