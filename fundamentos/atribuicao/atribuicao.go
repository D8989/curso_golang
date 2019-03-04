package main

import "fmt"

func main() {
	var b byte = 3 // classico
	fmt.Println(b)

	i := 3 // inferencia
	fmt.Println(i)

	i += 3
	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x // troca o valor entre as variaveis
	fmt.Println(x, y)
}
