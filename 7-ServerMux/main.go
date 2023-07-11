package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)

	//Ao chamar a struct ele chama o método automáticamente
	mux.Handle("/blog", blog{title: "My blog"})

	http.ListenAndServe("8080", mux)

	mux2 := http.NewServeMux()

	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Marcio"))
	})
	http.ListenAndServe("8081", mux2)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
