package main

import (
	"fmt"
)

func main(){
	fmt.Println(sum(2,1))
	fmt.Printf("O tipo do retorno de sum(2,1) é %T\n", sum(2,1))

	fmt.Println(sum_and_compare(2,1))
	fmt.Println(sum_and_compare(5,7))
}

//	Função que recebe dois valores inteiros, soma e retorna um resultado inteiro
func sum(a int, b int) int{
	return a + b
}

/*
	Função que recebe dois valores inteiros, faz uma comparação da soma desses dois valores, 
	caso essa condição for satisfeita ele retorna a soma, que também é um inteiro e um booleano "true", 
	caso contrário ele retorna a soma e o valor booleano "false".
*/
func sum_and_compare(a, b int) (int, bool){
	if(a + b >= 10){
		return a + b, true
	}
	return a + b, false
}
