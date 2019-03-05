package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta a quantidade de elementos

	for i, numero := range numeros { // 'range' retorna dois valores: o indice e o valor
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { // ignora o indice dos elementos
		fmt.Println(num)
	}
}
