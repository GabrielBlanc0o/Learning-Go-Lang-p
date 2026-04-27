package main

import "fmt"

func main() {
	numeros := []int{10, 20, 30, 40}

	fmt.Println("=== Solo índice ===")
	for i := range numeros {
		fmt.Println(i)
	}

	fmt.Println("=== Solo valor ===")
	for _, v := range numeros {
		fmt.Println(v)
	}

	fmt.Println("=== Índice y valor ===")
	for i, v := range numeros {
		fmt.Printf("Índice %d = %d\n", i, v)
	}
}
