package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) // cast

	nota := 6.9
	notaFinal := int(nota) // 6
	fmt.Println(notaFinal)

	// cuidado...
	// fmt.Println("Teste " + string(97))
	// fmt.Sprint(97)

	// int para string
	fmt.Println("Teste ", strconv.Itoa(97))

	// string para int
	num, _ := strconv.Atoi("97")
	fmt.Println("Teste ", num)

	b, _ := strconv.ParseBool("True") // apenas 'true' e '1' são considerados verdadeiros
	if b {
		fmt.Println("Verdadeiro")
	}
}
