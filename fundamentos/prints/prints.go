package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// fmt.Println("O valor de x é " + x) // isso não é permitido

	// Solução 1
	xs := fmt.Sprint(x) // retorna o valor de x em string
	fmt.Println("O valor de x é " + xs + " unidades.")

	// Solução 2
	fmt.Println("O valor de x é", x, "unidades.")

	// Solução 3
	fmt.Printf("O valor é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	fmt.Printf("\n%v %v %v %v", a, b, c, d) // '%v' funciona com a maiorida dos tipos
}
