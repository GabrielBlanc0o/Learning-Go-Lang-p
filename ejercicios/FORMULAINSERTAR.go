package main

import "fmt"

func main() {

	ejemplo()
	ejercicio21()
	ejercicio22()
	ejercicio23()
	ejercicio24()
}

func ejemplo() {
	// slice de valores
	numeros := []int{1, 2, 3, 4, 6}
	fmt.Println(numeros)

	posicion := 4
	valor := 5

	numeros = append(numeros[:posicion], append([]int{valor}, numeros[posicion:]...)...)

	fmt.Println(numeros)
}

func ejercicio21() {
	nums := []int{20, 30, 40, 50}
	posicion := 0
	valorr := 10

	nums = append(nums[:posicion], append([]int{valorr}, nums[posicion:]...)...)
	fmt.Println(nums)
}

func ejercicio22() {
	nums := []int{10, 20, 30, 40}
	valor := 50

	nums = append(nums, valor)
	fmt.Println(nums)

}

func ejercicio23() {
	nums := []int{10, 20, 50, 60}
	posicion := 2
	valores := []int{30, 40}

	nums = append(nums[:posicion], append([]int{valores[0], valores[1]}, nums[posicion:]...)...)

	fmt.Println(nums)
}

func ejercicio24() {

	nums := []int{1, 2, 3, 4, 5}
	posicion := 2
	valor := 99

	for i := 0; i < 5; i++ {
		nums = append(nums[:posicion], append([]int{valor}, nums[posicion:]...)...)
		posicion++
	}

	fmt.Println(nums)

}
