package main

func somaInt(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func somaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

// Função generica:

// ----constrainst-----

type MyNumber int

type Number interface {
	~int | ~float64 

	/* O símbolo ~ é usado para indicar que o tipo pode ser um tipo definido pelo usuário que tem a mesma representação subjacente do tipo especificado na generic 	(neste caso, int ou float64). */
}

func soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Arthur": 10, "Maria": 20, "João": 30}
	m2 := map[string]float64{"Arthur": 10.5, "Maria": 20.5, "João": 30.5}
	m3 := map[string]MyNumber{"Arthur": 10, "Maria": 20, "João": 33, "Matheus": 47}
	println(somaInt(m))
	println(somaFloat(m2))
	println(soma(m3))
	println(compare(10, 10.0))
}
