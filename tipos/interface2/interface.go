package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func ligar(e esportivo) {
	e.ligarTurbo()
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"F50", false, 0}
	carro2.ligarTurbo()

	carro3 := ferrari{"F60", false, 0}
	ligar(&carro3)

	fmt.Println(carro1, carro2, carro3)
}
