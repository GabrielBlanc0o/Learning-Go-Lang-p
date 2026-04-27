package main

import "fmt"

func main() {
	fmt.Print("Lista: \n")
	original := []int{100, 200, 300, 400, 500}
	converted := []int{}
	converted = append(converted, original[0], original[len(original)-1])
	fmt.Print(converted, "\n")
}
