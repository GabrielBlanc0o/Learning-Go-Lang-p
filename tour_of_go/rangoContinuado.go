package main

import "fmt"

func main() {

	pow := make([]int, 10)
	// esto es los numeros de 0 a 10

	for i := range pow { // para i en el rango de pow sin valor solo indice index	z

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

/* ANALIZANDO esto es mejor dicho q si tenemos un espacio de 10 posiciones
movemos  un 1 en ese slice de puros 0 PERO cada vez q movemos un 1 a la
izquierda el valor es el mismo pero su duplicado

en binario por ejemplo
0000000001 es 1
0000000010 es 2
0000000100 es 4

por eso da esa ilusion de potencia

ya que la linea 12 es lo q se encarga y el limite de movimientos es dado por la
condicion de range


*/
