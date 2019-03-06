package main

import "fmt"

func main() {
	// var aprovados map[int]string // var <nome> map[<chave>]<valor>
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	// para adcionar elementos no map
	aprovados[123456789] = "Maria"
	aprovados[987654321] = "Pedro"
	aprovados[789456123] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF:%d)\n", nome, cpf)
	}

	// para ler
	fmt.Println(aprovados[123456789])

	// para excluir
	delete(aprovados, 987654321)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF:%d)\n", nome, cpf)
	}
}
