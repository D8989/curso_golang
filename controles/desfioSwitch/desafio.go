package main

import "fmt"

func notaParaConceito(n float64) string {
	nota := int(n)
	var conceito string
	switch {
	case nota == 10:
		fallthrough
	case nota >= 9:
		conceito = "A"
	case nota >= 7:
		conceito = "B"
	case nota >= 5:
		conceito = "C"
	case nota >= 3:
		conceito = "D"
	case nota >= 0:
		conceito = "E"
	default:
		conceito = "Nota invalida"
	}
	return conceito
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.8))
	fmt.Println(notaParaConceito(2.1))
}
