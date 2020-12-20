package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados

	aprovados := make(map[int]string)
	aprovados[1234567878] = "Maria"
	aprovados[9876543214] = "Pedro"
	aprovados[6574839201] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[9876543214])
	delete(aprovados, 1234567878)
	fmt.Println(aprovados)
}
