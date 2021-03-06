package main

import (
	"fmt"
	"time"

	myhtml "github.com/D8989/my-html"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := myhtml.Titulo(url1)
	c2 := myhtml.Titulo(url2)
	c3 := myhtml.Titulo(url3)

	// select - um switch especifico para a concorrencia
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.amazon.com",
		"https://www.youtube.com",
		"https://www.google.com",
	)

	fmt.Println(campeao)
}
