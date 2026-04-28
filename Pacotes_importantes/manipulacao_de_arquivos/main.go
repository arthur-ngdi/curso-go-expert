package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
/*
	f, err := os.Create("teste.txt")
	if err != nil {
		panic(err)
	}

	// length, err := f.WriteString("Manipulação de arquivos em Go\n")
	length, err := f.Write([]byte("Manipulação de arquivos em Go\n"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso, número de bytes escritos: %d\n", length)

	f.Close()

	// LEITURA
	arquivo, err := os.ReadFile("teste.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(arquivo)
	fmt.Println(string(arquivo))
*/
	// LEITURA EM PARTES
	f2, err := os.Open("teste.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f2)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
		fmt.Printf("Número de bytes lidos: %d\n", n)
	}

	//err = os.remove("teste.txt")
	/* if err != nil {
		panic(err)
	} */
}
