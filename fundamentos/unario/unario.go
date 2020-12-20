package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix
	x++
	y--

	fmt.Println(x, y)

	// a := x == y-- // nao pode utilizar unario dentro de uma expressao
}
