package main

import (
	"curso-go/matematica"
	"fmt"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v\n", s)
	fmt.Println(matematica.A)
	fmt.Println(matematica.B)
	c := matematica.Carro{
		Marca: "Toyota",
		Ano:   2020,
	}
	fmt.Println(c)
	c.Andar()

}
