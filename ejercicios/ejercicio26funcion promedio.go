package main

import "fmt"

func promedio(notas []float64) float64 {

	var promedio float64
	var cantidadNumeros float64
	for i := 0; i < len(notas); i++ {
		promedio += notas[i]
		cantidadNumeros++
	}
	if cantidadNumeros == 0 {
		return 0
	}
	return promedio / cantidadNumeros

}

func main() {
	notas := []float64{7.5, 8.0, 9.0, 6.5}
	fmt.Printf("Promedio: %.2f\n", promedio(notas))
}
