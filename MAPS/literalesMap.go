package main

import "fmt"

type Ubicacion struct {
	Lat, Long float64
}

var m = map[string]Ubicacion{

	"Bell Labs": Ubicacion{
		40.68433, -74.39967,
	},
	"Google": Ubicacion{
		37.42202, -122.08408,
	},
}

func main() {

	fmt.Println(m)

}
