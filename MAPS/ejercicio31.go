package main

import "fmt"

func main() {
	ejercicio31()
	ejercicio32()
	ejercicio33()
}

func ejercicio31() {

	calificaciones := map[string]int{
		"Ana":    85,
		"Luis":   92,
		"Carlos": 78,
	}

	// Mostrar la calificación de Ana
	fmt.Println("Calificación de Ana:", calificaciones["Ana"])
}

func ejercicio32() {

	calificaciones := map[string]int{
		"Ana":  85,
		"Luis": 92,
	}

	calificaciones["Carlos"] = 88
	calificaciones["Ana"] = 90

	// Modificar la calificación de "Ana" a 90

	fmt.Println(calificaciones)

}

func ejercicio33() {

	calificaciones := map[string]int{
		"Ana":  85,
		"Luis": 92,
	}

	nombre := "Carlos"

	// Verificar si "Carlos" existe
	// Pista: valor, existe := map[clave]

	nota, existe := calificaciones[nombre]
	if existe {
		fmt.Println("Calificación de", nombre, ":", nota)
	} else {
		fmt.Println(nombre, "no encontrado")
	}

}
