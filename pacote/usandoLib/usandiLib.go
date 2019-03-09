package main

// comando para instalar pelo github
// go get -u github.com/<PATH>

// O ideal é realizar um 'git clone' na pasta go/src/github.com/ vazio
// e diretamente de lá criar seu proprio pacote

import (
	"fmt"

	"github.com/cod3rcursos/goarea"
)

func main() {
	fmt.Println(goarea.Circ(4.0))
}
