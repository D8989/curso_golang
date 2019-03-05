package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	s1[0] = 7
	fmt.Println(s1, s2)
	/*
		array := [6]int{5, 6, 7, 8, 9, 10}
		fmt.Println(array)

		s1 := array[:len(array)/2]
		s2 := array[len(array)/2:]

		fmt.Println(s1, s2)

		s1[0] = 0
		s2[0] = 0
		fmt.Println(array, s1, s2)
	*/
}
