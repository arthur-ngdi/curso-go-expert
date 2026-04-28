package main

import "fmt"

func main() {
	a := 10
	//fmt.Println(&a) //o "&" é o operador de endereço, ele retorna o endereço de memória da variável

	var ponteiro *int = &a
	*ponteiro = 20

	fmt.Println(a)

	b := &a
	*b = 30
	fmt.Println(a)
}
