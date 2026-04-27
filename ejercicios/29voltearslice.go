package main

import "fmt"

func invertir(slice []int) {
	for i := 0; i < len(slice)/2; i++ {
		var temp int
		temp = slice[i]
		slice[i] = slice[len(slice)-i-1]
		slice[len(slice)-i-1] = temp
	}
}

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println("Antes:", numeros) // [1 2 3 4 5]
	invertir(numeros)
	fmt.Println("Después:", numeros) // [5 4 3 2 1]
}
