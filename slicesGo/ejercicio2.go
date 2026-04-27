package main

import "fmt"

func main() {
	sumaElementos()
	numerosPares()
	existeElemento()
	eliminarElementoporIndice()
	invertir()
}

func sumaElementos() {

	numeros := []int{10, 20, 30, 40, 50}

	sumatotal := 0
	for _, t := range numeros {
		sumatotal += t
	}
	fmt.Println()
	fmt.Print("La suma es: ", sumatotal, "\n")
}

func numerosPares() {

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numerosPares := []int{}

	for _, n := range numeros {
		if n%2 == 0 {
			numerosPares = append(numerosPares, n)
		}
	}

	fmt.Println(numerosPares)
}

func existeElemento() {
	frutas := []string{"manzana", "pera", "uva", "naranja", "sandía"}
	buscado := "uva"

	for _, fruta := range frutas {
		if fruta == buscado {
			fmt.Println("Encontrado  ")
		}
	}
}

func eliminarElementoporIndice() {
	numeros := []int{10, 20, 30, 40, 50}

	numeros = append(numeros[:2], numeros[3:]...)

	fmt.Println(numeros)

}

func invertir() {

	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(numeros)

	for i := 0; i < len(numeros)/2; i++ {
		temp := numeros[i]
		numeros[i] = numeros[len(numeros)-i-1]
		numeros[len(numeros)-i-1] = temp
	}
	fmt.Println(numeros)

}
