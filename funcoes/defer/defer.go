package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("FIM!")
	if numero > 5000 {
		fmt.Println("valor alto...")
		return 5000
	} else {
		fmt.Println("valor baixo...")
		return numero
	}
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}