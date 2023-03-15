package main

import "fmt"

const a = "Hello, World!"

type ID int

// Variaveis de escopo global para todos os package main
var (
	b bool    = true
	c int     = 10
	d string  = "Wesley"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de F é %T", f)
}
