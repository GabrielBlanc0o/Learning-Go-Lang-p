package main

import "golang.org/x/tour/wc"

func WordCount(s string) map[string]int {

	result := make(map[string]int) //crear slice map clave valor string y enteros numeros
	word := ""                     // variable palabra vacia
	for _, r := range s {          // iteramos sobre el rango de s para ir palabra por palabra
		if r != ' ' { // el espacio es un separador si no es un espacio en blanco
			word = word + string(r) // agrega un espacio
		} else { // si no
			if word != "" { // si no es un espacio osea si hay palabra
				result[word]++ // meter word a clave valor map
				word = ""      // volver a dejar vacia para q reinicie la iteracion
			}
		}
	}
	//ultima palabra
	if word != "" { // si palabra no es un espacio vacio
		result[word]++ // q agregue al map de clave valor
	}
	return result // devuelva todo
}

func main() {
	wc.Test(WordCount) // usando la libreria de wc con test probar la logica de la
	//funcion de arriba y si pasa OK
}
