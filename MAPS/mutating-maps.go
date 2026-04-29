package main

import "fmt"

func main() {

	value := make(map[string]int) // creamos un slice de clave string valor int

	value["Answer"] = 42                         // creamos el primer de tipo answer en string y su valor en int
	fmt.Println("EL valor es:", value["Answer"]) // imprimimos

	value["Answer"] = 48                         // actualizamos el valor de la misma manera
	fmt.Println("EL valor es:", value["Answer"]) // imprimos de la misma maneraque arriba

	delete(value, "Answer") // usando el metodo delete() eliminamos todo lo que contenga la clave answer
	fmt.Println("EL valor es:", value["Answer"])

	v, ok := value["Answer"]                // para poder llamar cada elemento de cierta clave en especial creamos dos variables para poder llamar cada una de manera independiente
	fmt.Println("EL valor", v, "Esta?", ok) // imprimimos los dos valores declarados sobre el slice con tipo de clave valor arriba
}
