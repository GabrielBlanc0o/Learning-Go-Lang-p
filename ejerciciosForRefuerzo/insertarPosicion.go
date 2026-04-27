package main

import "fmt"

func main() {

	nums := []int{10, 20, 40, 50}
	posicion := 2 // después del 20
	valor := 30
	numsnew := []int{}
	
	numsnew = append(nums[:posicion:posicion], valor)
	nums = append(numsnew, nums[posicion:]...)

	fmt.Println(nums)

}
