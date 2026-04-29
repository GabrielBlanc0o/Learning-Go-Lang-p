package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}

	gestorNumeros(slice)
	filtrarNumeros(slice)
	modificarValores(slice)
}
func gestorNumeros(slice []int) {

	fmt.Println(slice)

	var largo int
	var capacidad int
	largo = len(slice)
	capacidad = cap(slice)

	fmt.Printf("Largo : %d , Capacidad: %d", largo, capacidad)
	fmt.Println()

}

func filtrarNumeros(slice []int) []int {

	var numerosMayores []int

	for _, numero := range slice {
		if numero > 25 {
			numerosMayores = append(numerosMayores, numero)
		}
	}
	fmt.Println("Numeros mayores a 25 :")
	return numerosMayores
}

func modificarValores(slice []int) {

	for i := 0; i < len(slice); i++ {

		slice[i] = slice[i] * 2

	}
	fmt.Println(slice)

}
