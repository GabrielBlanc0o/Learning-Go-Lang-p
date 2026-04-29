package main

import "fmt"

func main() {

	var s []int // slice vacio llamado s no tiene largo ni capacidad
	imprimirSlice(s)

	s = append(s, 0) // el append funcional cuando no hay nada o se usa nil
	imprimirSlice(s) // le agregamos un valor 0 capacidad y largo 1

	// lo hacemos crecer al slice como sea necesario
	s = append(s, 1) // ahora tiene 0 y 1 osea dos de largo y dos de capacidad
	imprimirSlice(s)

	s = append(s, 2, 3, 4) // podemos agregar mas de un elemento al tiempo
	imprimirSlice(s)       // tiene 0 1 2 3 4 5 de largo y 6 de capacidad porq queda un espacio ok po rel s

}

func imprimirSlice(s []int) {

	fmt.Printf("len=%d cap%d %v\n", len(s), cap(s), s)

}
