package main

import (
	"fmt"
)

func main() {
	//trabalhando com hashtables

	salarios := map[string]int{
		"João": 3000,
		"Maria": 4000,
		"Pedro": 3500,
	}
	//fmt.Println(salarios["João"])
	delete(salarios, "João")
	//fmt.Println(salarios)

	// sal := make(map[string]int)

	//for retornando chave e valor do map
	for nome, salario := range salarios{
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	//for com black identifier
	for _ , salario := range salarios{
		fmt.Printf("O salário é %d\n", salario)
	}
}

