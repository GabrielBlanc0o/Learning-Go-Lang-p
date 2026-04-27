package main

import "fmt"

func mainc() {
	fmt.Println("¡Go funcionando en Linux Mint!")

	// Esto es como vector<int> en C++
	numeros := []int{5, 10, 15, 20, 25}

	fmt.Println("Lista de números:", numeros)
	fmt.Println("Cantidad de elementos:", len(numeros))

	// Suma como en C++
	suma := 0
	for i := 0; i < len(numeros); i++ {
		suma = suma + numeros[i]
	}
	fmt.Println("Suma total:", suma)
}
