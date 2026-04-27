package main

import "fmt"

func main() {

	original := []int{10, 20, 30, 40, 50, 60}
	converted := []int{}
	converted = original[1 : len(original)-1]
	fmt.Print(converted, "\n")

}
