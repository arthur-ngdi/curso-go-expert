package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Person interface {
	DeativateClient()
}

type Client struct {
	Name    string
	Age     int
	Status  bool
	Address Address
}

func (c Client) DeactivateClient() {
	c.Status = false
	fmt.Printf("O cliente %s foi desativado\n", c.Name)
}

func main() {
	c1 := Client{
		Name:   "Arthur",
		Age:    25,
		Status: true,
	}

	fmt.Printf("O nome do cliente é %s, ele tem %d anos e seu status é %t\n", c1.Name, c1.Age, c1.Status)
	fmt.Println(c1.Name)

	//c1.Status = false
	//fmt.Printf("O nome do cliente é %s, ele tem %d anos e seu status é %t\n", c1.Name, c1.Age, c1.Status)

	c1.Address.City = "Santa Rita Do Sapucaí"
	fmt.Printf("O nome do cliente é %s, ele tem %d anos, seu status é %t e mora na cidade de %s\n", c1.Name, c1.Age, c1.Status, c1.Address.City)

	c1.DeactivateClient()
	fmt.Printf("O nome do cliente %s foi desativado\n", c1.Name)
}
