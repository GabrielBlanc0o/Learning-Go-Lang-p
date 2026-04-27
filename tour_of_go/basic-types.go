/*

	|--TIPOS DE DATOS BASICOS---|

	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // alias for uint8

	rune // alias for int32
		// represents a Unicode code point

	float32 float64

	complex64 complex128

*/

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false                // tipo bool
	MaxInt uint64     = 1<<64 - 1            // valor amplios de 64 bits para int
	z      complex128 = cmplx.Sqrt(-5 + 12i) //numero complejos para este caso raiz cuadrada
)

func main() {

	fmt.Printf("Tipo: %T Valor: %v\n", ToBe, ToBe)
	fmt.Printf("Tipo: %T Valor: %v\n", MaxInt, MaxInt)
	fmt.Printf("Tipo: %T Valor: %v\n", z, z)

}
