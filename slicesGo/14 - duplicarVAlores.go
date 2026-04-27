package main

import "fmt"

func main() {

	numeros := []int{2, 4, 6, 8, 10}

	for i := 0; i < len(numeros); i++ {
		numeros[i] = numeros[i] * 2
	}
	fmt.Print(numeros)
}
