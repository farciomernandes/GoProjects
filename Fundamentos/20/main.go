package main

type MyNumber int

// Constraint, é um tipo para falar que pode ser um valor ou outro
type NumberConstraint interface {
	~int | ~float64
	//O "~" avisa para que uma variavel do tipo INT e FLOAT64 pode ser considerado "NumberConstraint"
}

// Esse T pode ser do tipo INT ou FLOAT64
func Soma[T NumberConstraint](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// A "comparable" permite ter mais de um valor, porém os valores devem ser comparaveis um com o outro. Ou seja, "a" e "b" devem ter tipos
// que possam ser comparados! ( compara em relação a tipo, não funciona para comparar tamanhos)
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Wesley": 1000.20, "João": 2000.200, "Maria": 3000.20}
	m3 := map[string]MyNumber{"Wesley": 1000, "João": 2000, "Maria": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))
}
