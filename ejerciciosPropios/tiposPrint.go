package main

import "fmt"

func main() {

	edad := 18
	nombre := "gabito"

	//print basico
	fmt.Println("Hola negros")
	fmt.Println("mi edad es", edad, "me llamo", nombre)
	fmt.Printf("mi edad es %v me llamo %v ", edad, nombre)

	// cadenas guardadas y formateadas

	var str = fmt.Sprintf("mi edad es %v me llamo %v ", edad, nombre)

	fmt.Println("\nresultado de cadena es: ", str)
}
