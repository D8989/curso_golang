package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // pegando o endereco de i

	*p++
	i++

	// Go nao tem aritmetica de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}
