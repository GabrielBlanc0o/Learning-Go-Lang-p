package main

import "fmt"

func main() {

	a := make([]int, 5)
	imprimirSlice("a", a)

	b := make([]int, 0, 5)
	imprimirSlice("b", b)

	c := b[:2] // toma el primer 0 y 1 de b
	imprimirSlice("c", c)

	d := c[2:5] // toma de 2 a 4 de c q son 0 y 1  y 2
	imprimirSlice("d", d)

}

func imprimirSlice(s string, x []int) {

	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)

	// slice original su largo y capacidad y x es su slice final segun la propiedad

}
