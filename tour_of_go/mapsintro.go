package main

import "fmt"

// creamos una estructura para datos decimales
type Coordenadas struct {
	Lat, Long float64
}

// variable llamada m de mapa q contenga un string y la estructura con los dos flotantes
var m map[string]Coordenadas

func main() {

	m = make(map[string]Coordenadas) // m es un slice con el map de arriba igual
	m["Bell Labs"] = Coordenadas{    // llenamos m con los dos datos el string y coordenadas
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"]) // imprime m con el string dado y el dato especial q son coordenandas

}
