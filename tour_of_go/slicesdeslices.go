package main

import (
	"fmt"
	"strings"
)

func main() {

	//creamos la tabla de tic tac toe

	board := [][]string{

		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// los jugadores toman los turnos

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	// bucle para imprimir todo con movimientos

	for i := 0; i < len(board); i++ {

		fmt.Printf("%s\n", strings.Join(board[i], " "))

	}

}
