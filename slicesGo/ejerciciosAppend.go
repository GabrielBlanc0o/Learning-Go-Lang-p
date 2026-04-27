package main

import "fmt"

func main() {
	ejercicio6()
	ejercicio7()
	ejercicio8()
	ejercicio9()
}

func ejercicio6() {

	numeros1 := []int{1, 2, 3}

	numeros1 = append(numeros1, 4, 5, 6)
	fmt.Println()
	fmt.Print(numeros1)
	fmt.Println()
}

func ejercicio7() {

	slice := []int{}

	for i := 1; i < 11; i++ {
		slice = append(slice, i)
	}
	fmt.Print(slice)

}

func ejercicio8() {
	numeros1 := []int{1, 2, 3}
	numeros2 := []int{4, 5, 6}

	numeros1 = append(numeros1, numeros2...)
	fmt.Println()
	fmt.Print(numeros1)
	fmt.Println()
}

func ejercicio9() {

	numeros := []int{5, 12, 7, 3, 18, 9, 21, 4}
	mayores := []int{}

	for _, value := range numeros {
		if value > 10 {
			mayores = append(mayores, value)
		}
	}
	fmt.Print(mayores)

}
