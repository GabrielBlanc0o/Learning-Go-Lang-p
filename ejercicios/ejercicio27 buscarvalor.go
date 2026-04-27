package main

import "fmt"

func existe(slice []int, buscado int) bool {

	var encontrado bool

	for _, num := range slice {
		if num == buscado {
			encontrado = true
			break
		}
	}
	return encontrado
}

func main() {
	numeros := []int{10, 20, 30, 40, 50}
	fmt.Println("Lista: ", numeros)
	fmt.Println("¿Existe 30?", existe(numeros, 30))
	fmt.Println("¿Existe 99?", existe(numeros, 99))
}
