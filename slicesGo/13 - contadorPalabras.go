package main

import "fmt"

func main() {

	palabras := []string{"gato", "elefante", "sol", "mariposa", "casa", "computadora"}
	var contador int

	for _, palabra := range palabras {
		if len(palabra) > 5 {
			contador++
		}
	}
	fmt.Println("Palabras con mas de 5 letras:", contador)

}
