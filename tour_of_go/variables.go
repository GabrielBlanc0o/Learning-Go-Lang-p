package main

import "fmt"

func main() {
	var i, j int = 1, 2
	variables()
	variablesInicializadas(i, j)

}

func variables() {
	var c, python, java bool
	var i int
	fmt.Println(i, c, python, java)

}

func variablesInicializadas(x, y int) (int, int) {

	var c, python, java = true, false, "no!"
	fmt.Println(x, y, c, python, java)
	return x, y

}

// variables con inicializacion
