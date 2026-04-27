package main

import "fmt"

func concatenar(a, b, c []int) []int {
	a = append(a[:], append(b, c...)...)
	return a
}

func main() {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	z := []int{7, 8, 9}

	resultado := concatenar(x, y, z)
	fmt.Println(resultado) // [1 2 3 4 5 6 7 8 9]
}
