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
	//criando e percorrendo arrays
	var g [3]int
	g[0] = 1
	g[1] = 2
	g[2] = 3

	fmt.Println(len(g)-1)
	fmt.Println(g[len(g)-1])

	for i, v := range g{
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}
}

func x(){
}
