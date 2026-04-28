package main

import "fmt"

func showType(t interface{}) {
	fmt.Printf("To tipo da variável é: %T\n e o valor é: %v\n", t, t)
}

func main() {
	var x interface{} = 10
	var y interface{} = "Hello"

	showType(x)
	showType(y)
}
