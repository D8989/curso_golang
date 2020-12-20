package main

import "fmt"

// Nao existe o ternario em Go
func opterResultado(nota float64) string {
	if nota >= 6.0 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(opterResultado(6.2))
}
