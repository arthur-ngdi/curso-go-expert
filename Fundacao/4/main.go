package main

import (
	"fmt"
)

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
	fmt.Printf("O tipo de A é %T\n", a)
	fmt.Printf("O tipo de B é %T\n", b)
	fmt.Printf("O tipo de C é %T\n", c)
	fmt.Printf("O tipo de D é %T\n", d)
	fmt.Printf("O tipo de E é %T\n", e)
	fmt.Printf("O tipo de F é %T\n", f)
}

func x(){
}
