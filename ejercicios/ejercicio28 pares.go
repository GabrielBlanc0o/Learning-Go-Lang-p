package main

import "fmt"

func filtrarPares(slice []int) []int {
	pares := []int{}
	for _, nums := range slice {
		if nums%2 == 0 {
			pares = append(pares, nums)
		}
	}
	return pares
}

func main() {
	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pares := filtrarPares(original)

	fmt.Println("Original:", original) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println("Pares:", pares)       // [2 4 6 8 10]
}
