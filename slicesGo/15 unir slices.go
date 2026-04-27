package main

import "fmt"

func main() {

	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8}
	resultado := []int{}

	for i := 0; i < len(a); i++ {
		resultado = append(resultado, a[i], b[i])
	}

	fmt.Println(resultado)
}
