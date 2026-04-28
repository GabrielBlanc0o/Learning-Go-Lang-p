package main

import (
	"fmt"
	"math"
)

func main() {

	var x, y int = 3, 4
	var f = math.Sqrt(float64(x*x + y*y)) // raiz cuadrada de 9 + 16 = 25 =  5
	var z uint = uint(f)
	fmt.Println(x, y, z) // 3 , 4 , 5

}
