package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128} // potencias de dos

func main() {

	for i, v := range pow {

		fmt.Printf("2**%d = %d\n", i, v)
		/*i incia en 0 y v q es el rango de el slice de arriba seria
		la copía del elemento de ese index osea
		si i es 0  v seria el primero elemento de pow
		osea si i es 0 v es 1
		i = 1 v = 2
		i = 2 v = 4

		la ilusion de potencia es porque el print se le pone 2** pero esto
		solo es visual
		*/
	}

}
