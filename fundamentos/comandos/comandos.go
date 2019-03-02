package main

import "fmt"

func main() {
	fmt.Printf("Outro programa em %s!!!!!!\n", "Go")
}

// go vet <arquivo.go> -> verifica erros no arquivo .go
// go build <arquivo.go> -> cria um executavel
// go run <arquivo.go> -> compila e roda o programa
// go get -u github.com/PATH -> instala um pacote Go do github
