package main

import "fmt"

// O 'init' é exécutado antes do main

func init() {
	fmt.Println("inicializando...")
}

func main() {
	fmt.Println("Main...")
}
