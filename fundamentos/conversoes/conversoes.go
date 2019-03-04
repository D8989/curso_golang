package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado
	fmt.Println("Teste " + string(97)) // -> ele tarnsforma o número pela tabela ASCII

	// int para string
	fmt.Println("Teste " + strconv.Itoa(97))

	// string para int
	num, _ := strconv.Atoi("97")
	fmt.Println(num - 96)

	b, _ := strconv.ParseBool("true") // true == 1
	if b {
		fmt.Println("Verdadeiro")
	}
}
