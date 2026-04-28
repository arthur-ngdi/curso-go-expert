package main

import (
	"fmt"
)

func main(){
	dobro := func () int{
		return sum(1,2,3,4,5,6,7,8,9,10) * 2
	}

	fmt.Println(dobro)
}

//	Função que recebe n valores inteiros, soma e retorna um resultado inteiro
func sum(numbers ...int) int{
	total := 0
	for _, number := range numbers{
		total += number
	}
	return total
}