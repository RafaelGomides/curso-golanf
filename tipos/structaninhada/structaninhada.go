package main

import "fmt"

type Item struct {
	idItem     int
	Quantidade int
	Preco      float64
}

type pedido struct {
	idPedido int64
	itens    []Item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.Preco * float64(item.Quantidade)
	}
	return total
}

func main() {
	pedido := pedido{
		idPedido: 1,
		itens: []Item{
			Item{1, 2, 12.10},
			Item{2, 1, 23.49},
			Item{11, 100, 1.23},
		},
	}
	fmt.Println(pedido.valorTotal())
}
