package main

import "fmt"

func add(x, y int) int {

	return x + y

}

func swap(x, y string) (string, string) { // recibe dos string y envia dos
	return y, x // orden contrario

}

func main() {

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world") // declara valores
	fmt.Println(a, b)              // imprime

}
