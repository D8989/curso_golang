package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!!") // executado logo antes do retorno
	if numero > 5000 {
		fmt.Println("Valor Alto...")
		return 5000
	}

	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
