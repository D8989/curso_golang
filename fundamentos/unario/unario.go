package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix
	x++ // x = x + 1
	fmt.Println(x)

	y-- // y = y - 1
	fmt.Println(y)

	// fmt.Print( x == y--) não é permitido fazer esta op dentro de uma expressão

	// ! -> negação lógica
}
