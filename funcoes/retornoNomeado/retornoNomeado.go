package main

import "fmt"

func troca(p1, p2 int) (segundo int, primeiro int) { // 'segundo' e 'primeiro' são variaveis que pertencem ao escopo da função
	segundo = p2
	primeiro = p1
	return // retorno limpo
}

// retorna na ordem que declara o "retorno nomeado"
func main() {
	x, y := troca(2, 3)
	fmt.Println(x, y)
}
