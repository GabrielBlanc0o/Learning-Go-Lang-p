package main

import "fmt"

func main() {

	crearAgregar()
	acceder()
	recorrer()

}

func crearAgregar() {

	// Crear slice vacío de enteros
	var numeros []int

	fmt.Println("Slice vacío:", numeros)
	fmt.Println("Cantidad:", len(numeros)) // cantidad de elementos

	// Agregar elementos
	numeros = append(numeros, 5)
	numeros = append(numeros, 10)
	numeros = append(numeros, 15)

	fmt.Println("Slice después de agregar:", numeros)
	fmt.Println("Cantidad ahora:", len(numeros))
	fmt.Println()
}

func acceder() {

	// Crear slice con valores
	colores := []string{"rojo", "verde", "azul"}

	// Acceder como en C++
	fmt.Println("Primer color:", colores[0])
	fmt.Println("Segundo color:", colores[1])
	fmt.Println("Tercer color:", colores[2])

	// Modificar un elemento
	colores[1] = "amarillo"
	fmt.Println("Después de modificar:", colores)
	fmt.Println()
}

func recorrer() {

	frutas := []string{"manzana", "pera", "uva", "naranja"}

	// Forma 1: solo valores
	fmt.Println("=== Solo valores ===")
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}

	// Forma 2: índice y valor
	fmt.Println("\n=== Índice y valor ===")
	for i, fruta := range frutas {
		fmt.Printf("Posición %d: %s\n", i, fruta)
	}
}
