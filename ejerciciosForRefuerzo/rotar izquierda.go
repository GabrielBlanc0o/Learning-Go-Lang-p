package main

import "fmt"

func main() {

	nums := []int{10, 20, 30, 40, 50}
	n := 2
	nums = append(nums[n:], nums[:n]...)
	fmt.Println(nums)

}
