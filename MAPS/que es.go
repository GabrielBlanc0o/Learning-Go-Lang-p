package main

import "fmt"

/*
es una coleccion que guarda pares clave - valor
tipo diccionario
por ejemplo clave puede ser nombre y el valor gabriel
{nombre:Gabriel}
*/
// map vacio
func main() {
	edades := make(map[string]int)

	edades["Ana"] = 25
	edades["Luis"] = 30
	edades["Carlos"] = 28

	// obtener un valor
	fmt.Println("Usuarios", edades)
	fmt.Println(edades["Ana"]) //25

	//eliminar
	delete(edades, "Luis")
	fmt.Println("Usuarios actualizados", edades)

	// cuantos elementos hay
	fmt.Println(len(edades))

}
