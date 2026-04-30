package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	resultado := make(map[string]int)
	palabra := ""

	for _, r := range s {
		if r != ' ' {
			palabra += string(r)
		} else {
			if palabra != "" {
				resultado[palabra]++
				palabra = ""
			}

		}
	}

	if palabra != "" {
		resultado[palabra]++
	}

	return resultado

}

func main() {
	wc.Test(WordCount) // usando la libreria de wc con test probar la logica de la
	//funcion de arriba y si pasa OK
}
