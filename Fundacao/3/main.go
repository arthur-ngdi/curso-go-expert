package main

const a = "Hello, World!"

//declarando meu próprio tipo de variavel
type ID int

var(
	b bool // assume "false" por padrão
	c int // assume o valor 0 por padrão
	d string //assume vazio por padrão
	e float64 //assume o valor 0.0 por padrão

	f ID
)

func main() {
	println(b)
	println(c)
	println(d)
	println(e)

	f = 10
	println(f)
}

func x(){
}
