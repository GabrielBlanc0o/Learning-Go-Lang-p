package main

import "fmt"

// Escribí una función que reciba un slice, una posición y un valor,
// y devuelva un NUEVO slice con el valor insertado.
// NO modificar el slice original.

func insertar(slice []int, posicion int, valor int) []int {
	nuevo := []int{}
	nuevo = append(nuevo, slice[:posicion]...)
	nuevo = append(nuevo, valor)
	nuevo = append(nuevo, slice[posicion:]...)
	return nuevo
}

func main() {
	original := []int{10, 20, 40, 50}
	nuevo := insertar(original, 2, 30)

	fmt.Println("Original:", original) // [10 20 40 50]
	fmt.Println("Nuevo:", nuevo)       // [10 20 30 40 50]
}
