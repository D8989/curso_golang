package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operacao bloqueande
	fmt.Println("So depois que for lido")
}

func main() {
	c := make(chan int)

	go rotina(c)

	fmt.Println(<-c) // operacao bloqueante
	fmt.Println("foi lido")
	fmt.Println(<-c) // deadlock
}
