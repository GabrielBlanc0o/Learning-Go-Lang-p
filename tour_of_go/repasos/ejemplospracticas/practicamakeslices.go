package main

import "fmt"

func main() {

	frutas := make([]string, 0, 10)
	frutas = append(frutas, "Manzana", "Pera", "Banana")
	frutas = append(frutas, "Sandía", "Mango", "Papaya")
	soloDos := frutas[1:3]
	extendido := soloDos[:cap(soloDos)]
	extendido[0] = "Misterio"
	extendido[4] = "Misterio"

	fmt.Printf("largo %d , capacidad %d, contenido %v", len(frutas), cap(frutas), frutas)
}
