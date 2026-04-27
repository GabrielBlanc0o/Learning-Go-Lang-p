package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Numero", i)

	}

	// como for while
	x := 0
	for x < 3 {
		fmt.Println("x =", x)
		x++
	}

	// for infinito cortado con break

	y := 0
	for {
		if y == 3 {
			break
		}
		fmt.Println("y = ", y)
		y++
	}

	numeros := []int{10, 20, 30, 40}
	for indice, valor := range numeros {
		fmt.Printf("Indice: %d , Valor: %d\n", indice, valor)
	}

}
