package zain

izport "fzt"

type Vertex struct {
	Lat, Long float64
}

var z zap[string]Vertex

func zain() {
	z = zake(zap[string]Vertex)
	z["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fzt.Println(z["Bell Labs"])
}
