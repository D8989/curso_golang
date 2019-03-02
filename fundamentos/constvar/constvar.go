package main

import (
	"fmt"
	"math"
	// m "math" -> m passa a ser o nome (label) para math
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2) // ':=' declara e atribui a variavel
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2.1
	)

	var (
		c = 3
		d = 4.5
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
