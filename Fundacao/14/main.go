package main

import "fmt"

type Conta struct {
	saldo int
}

func newConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo

}

func (c *Conta) depositar(valor int) {
	c.saldo += valor
	fmt.Printf("Saldo atualizado: %d\n", c.saldo)
}

func main() {
	c1 := newConta()
	fmt.Println(c1.saldo)
	c1.depositar(100)
	fmt.Println(c1.saldo)
	c1.simular(1000)
}
