package matematica

// para que a função, struct (e tudo o que está relacionado com a struct), variável, etc seja exportada, ou seja, possa ser utilizada fora do pacote, ela deve começar com letra maiúscula. Caso contrário, ela é privada ao pacote e não pode ser acessada de fora.

func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 10
var B int = 20

type Carro struct {
	Marca string
	Ano   int
}

func (c Carro) Andar() {
	println("carro andando")
}