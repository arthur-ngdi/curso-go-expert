package main

const a = "Hello, World!"

var(
	b bool // assume "false" por padrão
	c int // assume o valor 0 por padrão
	d string //assume vazio por padrão
	e float64 //assume o valor 0.0 por padrão
)

func main() {
	println(b)
	println(c)
	println(d)
	println(e)
}