package main

import "fmt"

type item struct {
	protudoID, qtde int
	// qtde      int
	preco float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{
				protudoID: 1,
				qtde:      2,
				preco:     12.10, // mais legivel
			},
			item{2, 1, 23.49},
			item{11, 100, 3.13},
		},
	}

	fmt.Printf("Valor total do pedido é R$ %.2f\n", pedido.valorTotal())
}
