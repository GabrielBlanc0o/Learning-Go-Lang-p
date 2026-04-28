package main

import "fmt"

func main() {
	nombres := [4]string{ // array de 4 nombres

		"Pepe",
		"Juan",
		"California",
		"Chong",
	}

	fmt.Println(nombres) // ["Pepe", "Juan","California","Chong"]

	a := nombres[0:2]
	b := nombres[1:3]
	fmt.Println(a, b) // [Pepe , Juan ] [Juan ,California]

	b[0] = "XXX"

	fmt.Println(a, b)    // [Pepe , XXX ] [XXX,California]
	fmt.Println(nombres) // ["Pepe", "XXX","California","Chong"]

}
