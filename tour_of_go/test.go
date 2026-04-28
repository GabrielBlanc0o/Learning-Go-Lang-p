package main

import (
	"fmt"
)

type Estructura struct {
	X, Y int
}

func mover(v *Estructura, dx int, dy int) {
	v.X += dx
	v.Y += dy
}

func ImprimirEstructura(v Estructura) {
	fmt.Println("X:", v.X, "Y:", v.Y)
}

func main() {

	v := Estructura{1, 2}
	ImprimirEstructura(v)
	mover(&v, 3, 4) // deberia quedar (4,6)
	ImprimirEstructura(v)

}
