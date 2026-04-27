package main

import "fmt"

func main() {

	nums := []int{3, 8, 5, 2, 9, 4, 7, 6}
	pares := []int{}
	impares := []int{}

	for _, v := range nums {
		if v%2 == 0 {
			pares = append(pares, v)
		} else if v%2 != 0 {
			impares = append(impares, v)
		}
	}
	fmt.Println(pares)
	fmt.Println(impares)
}
