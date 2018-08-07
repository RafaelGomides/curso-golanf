package main

import "fmt"

type produto struct {
	Nome     string
	Preco    float64
	desconto float64
}

// Método: função com receiver
func (p produto) precoComDesconto() float64 {
	return p.Preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		Nome:     "Teste",
		Preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Notebook", 520.00, 10.0}
	fmt.Println(produto2, produto2.precoComDesconto())
}
