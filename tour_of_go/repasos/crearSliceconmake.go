package main

import (
	"fmt"
)

func main() {

	a := make([]int, 5) // crear slice con make de tamaño 5
	imprimirSlice("a", a)

	b := make([]int, 0, 5) // crear slice con make de tamaño 0 a 5
	imprimirSlice("b", b)

	c := b[:2] // crear slice con make tomando el b de base q inicie en 0 a 2 (osea de 2 valores 0 y 1)
	imprimirSlice("c", c)

	d := c[2:5] // crear slice con make tomando el c de base q inicie en 2 a 5 (osea de 2 valores 2 y 4)
	imprimirSlice("d", d)
}

func imprimirSlice(s string, x []int) {

	fmt.Printf("%s len=%d cap=%d  %v\n", s, len(x), cap(x), x)

	/* este print lo q hace es imprimir el largo del slice dinamico creado,
	luego su capacidad   y la cadena final, s es el string en letra,
	len largo cap capacidad y x es el slice completo
	*/

}
