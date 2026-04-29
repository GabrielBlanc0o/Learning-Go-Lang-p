package main

import "fmt"

func main() {

	pow := make([]int, 10)
	// esto es los numeros de 0 a 10

	for i := range pow { // para i en el rango de pow sin valor solo indice index

		pow[i] = 1 << uint(i) // == 2**i
		// el primer valor iterado de i sobre el slice de pow serea igual a 1 y de tipo uint

	}

	for _, value := range pow { //SOLO valor iterable sin index
		// para imprimir el valor directamente del slice ya relleno no necesitamos
		//las posiciones cada una sera value en el rango del slice pow

		fmt.Printf("%d\n", value) // imprime cada una
		// 2**0 = 1 2**1 = 2 2**2 = 4 ... etc
	}
}
