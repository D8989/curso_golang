package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 12456.98,
			"Guga Periera":   123456.66,
		},
		"J": {
			"Jose silva": 9878.45,
		},
		"P": {
			"Pedro Junior": 1298.7,
			"Paula Marioa": 13456.65,
		},
	}

	fmt.Println(funcPorLetra)

	delete(funcPorLetra, "P")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)
	}
}
