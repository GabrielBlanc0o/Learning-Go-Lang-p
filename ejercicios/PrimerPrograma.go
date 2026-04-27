package main

import "fmt"

func main() {
	nota1 := 70
	nota2 := 85
	nota3 := 60

	fmt.Println("Tus notas son : ", nota1, nota2, nota3)

	p := calculaPromedio(nota1, nota2, nota3)
	va := verificarAprobacion(p)
	mensaje := mostrarMensajePersonalizado(p)

	fmt.Println("Promedio: ", p)
	fmt.Println("¿Aprobaste? ", va)
	fmt.Println(mensaje)

}

func calculaPromedio(a, b, c int) int {
	suma := a + b + c
	prom := suma / 3
	return prom
}

func verificarAprobacion(promedio int) bool {
	return promedio >= 60

}

func mostrarMensajePersonalizado(promedio int) string {
	if promedio >= 90 {
		return "Excelente"
	} else if promedio >= 70 {
		return "Bien"
	} else {
		return "Estudia mas"
	}
}
