package main

import "fmt"

// Não existe operador ternario no Go
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "aprovado"
	}

	return "reprovado"
}

func main() {
	fmt.Println(obterResultado(5.3))
}
