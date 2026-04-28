package main

/* a unica estrutura de loop em Go é o For
   
*/

func main() {

	for i := 0; i < 10; i++{
		print(i, ", ")
	}

	numeros := []string {"um", "dois", "tres", "quatro", "cinco"}

	println("\nLoop for range")
	for key, value := range numeros {
		println(key, value)
	}

	println("\nLoop for funcionando como while")
	i := 0
	for i < 10 {
		print(i, ", ")
		i++
	}

	println("\nLoop for funcionando como while infinito")
	//loop infinito
	//for{
	//	println("loop infinito")
	//}
}
