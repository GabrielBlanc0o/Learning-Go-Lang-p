package main

import "fmt"

func main() {
	// 1. Crear slice vacío de float64
	temperaturas := []float64{}

	// 2. Agregar 3 temperaturas
	temperaturas = append(temperaturas, 22.5)
	temperaturas = append(temperaturas, 12.3)
	temperaturas = append(temperaturas, 32.7)
	// agregá las otras dos...

	// 3. Encontrar la máxima
	max := temperaturas[0] // arrancás con la primera
	for _, t := range temperaturas {
		if t > max {
			max = t
		}
		// si t > max, entonces max = t
	}

	// 4. Calcular promedio
	suma := 0.0
	for _, t := range temperaturas {
		suma += t
		// suma += t
	}
	promedio := suma / float64(len(temperaturas))

	// 5. Mostrar resultados
	fmt.Println("Temperaturas:", temperaturas)
	fmt.Println("Máxima:", max)
	fmt.Println("Promedio:", promedio)
}
