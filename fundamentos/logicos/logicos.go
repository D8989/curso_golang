package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // ou exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv58, tv32, sorvete := compras(true, true)
	fmt.Printf("tv58: %t, tv32: %t, sorvete %t, saudável: %t", tv58, tv32, sorvete, !sorvete)
}
