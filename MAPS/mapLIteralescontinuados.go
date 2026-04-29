package main

import "fmt"

type UbicacionSitio struct {
	Lat, Long float64
}

var Lugares = map[string]UbicacionSitio{
	"Amazon": {40.68433, -74.39967},
	"Google": {37.42202, -122.08408},
}

func main() {

	fmt.Println(Lugares)
}
