package main

import "fmt"

/*
	// minha forma!
func fatorial(n int) (fat int, err error) {
	switch {
	case n < 0:
		fat, err = -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		fat, err = 1, nil
	default:
		fatorialInterior, _ := fatorial(n - 1)
		fat, err = n*fatorialInterior, nil
	}
	return
}
*/

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(4)
	fmt.Println(resultado)

	// resultado2 := fatorial(-4) -> O compilador lança o erro
	// fmt.Println(resultado2)
}
