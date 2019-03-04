package main

import "fmt"

func main() {
	i := 1

	// Go não existe aritmética de ponteiros
	var p *int

	p = &i // pegando o endereço fa variavel 'i'
	*p++   // desreferencia o ponteiro (pega o valor)
	i++

	// p++ -> isso não é permitido

	fmt.Println(p, *p, i, &i)
}
