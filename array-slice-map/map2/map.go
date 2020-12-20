package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jpse Joao":     123.78,
		"Gabriel Silva": 22234.98,
		"Pedor Junior":  77787.9,
	}

	funcsESalarios["Rafael Filho"] = 12345.00

	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
