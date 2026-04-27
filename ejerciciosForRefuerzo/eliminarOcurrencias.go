package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 2, 4, 2, 5}
	valorAEliminar := 2

	res := []int{}
	for _, v := range nums {
		if v != valorAEliminar {
			res = append(res, v)
		}
	}
	fmt.Print(res)

}
