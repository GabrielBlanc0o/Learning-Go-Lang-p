package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}

	ultimo := nums[len(nums)-1] //SIEMPRE ASI PARA ACCEDER AL ULTIMO VALOR DE UN SLICE

	for i := len(nums) - 1; i > 0; i-- {
		nums[i] = nums[i-1]

	}
	nums[0] = ultimo
	fmt.Print(nums)

}
