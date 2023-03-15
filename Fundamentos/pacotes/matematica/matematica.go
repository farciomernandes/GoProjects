package matematica

//Só é possível encontrar um método, função ou struct fora do pacote
//Se ela iniciar com inicial maiuscula

// Iniciar com maiscula = export default
func Soma[T int | float64](a, b T) T {
	return a + b
}
