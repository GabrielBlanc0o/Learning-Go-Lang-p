package main

import "fmt"

func main() {

	fmt.Print("Hola \n")
	tiposDatos()
}

// declarar tipos de datos

func tiposDatos() {
	var edad int = 18
	var nombre = "Gabo "
	edad2 := 18
	nombre2 := "Gabriel"
	fmt.Print(edad, edad2, nombre, nombre2)
}
