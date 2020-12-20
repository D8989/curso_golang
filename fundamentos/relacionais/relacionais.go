package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("strings:", "Banana" == "Banana")
	fmt.Println(3 != 2)
	fmt.Println(3 < 2)
	fmt.Println(3 > 2)
	fmt.Println(3 <= 2)
	fmt.Println(3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println(d1 == d2)
	fmt.Println(d1.Equal(d2))

	type Pessoa struct {
		nome string
	}

	p1 := Pessoa{"Joao"}
	p2 := Pessoa{"Joao"}
	fmt.Println(p1 == p2) // compara valor
}
