package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a * b)
	fmt.Println(a % b)

	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)

	c := 3.0
	d := 2.0

	fmt.Println(math.Max(c, d))
	fmt.Println(math.Min(c, d))
	fmt.Println(math.Pow(c, d))

}
