/*

los slices pueden ser creados con una funcion make

asi es como creamos arrays  de un tamaño dinamico

*/

// ejemplo:

package main

import (
	"fmt"
)

func main() {

	a := make([]int, 5)
	printSlice("a", a) // a ,5 ,5 [0,0,0,0,0]

	b := make([]int, 0, 5)
	printSlice("b", b) // b , 0 ,5 []

	c := b[:2]
	printSlice("c", c) // c ,2,5, [0,0]

	d := c[2:5]
	printSlice("d", d) // d , 3,3, [0,0,0]

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func gestorNumeros() {
	slice := make([]int, 5)
	slice = append(slice, 10, 20, 30, 40, 50)

	var largo int
	var capacidad int
	largo = len(slice)
	capacidad = cap(slice)

	fmt.Println("Largo : %d , Capacidad: %d", largo, capacidad)

}
