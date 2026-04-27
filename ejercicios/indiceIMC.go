package main

import "fmt"

func main() {

	peso, altura := preguntasUser()

	imc := formulaIMC(peso, altura)
	fmt.Println("Tu IMC es:", imc)

}

func preguntasUser() (float64, float64) {

	var peso float64
	var altura float64

	fmt.Println("Ingresa tu peso en kg: ")
	fmt.Scanln(&peso)
	fmt.Println("Ingresa tu altura en metros: ")
	fmt.Scanln(&altura)

	return peso, altura
}

func formulaIMC(peso float64, altura float64) float64 {

	var formula float64

	formula = peso / (altura * altura)

	fmt.Print("Tiene ")

	if formula < 18.5 {
		fmt.Print("bajo peso")
	} else if formula >= 18.5 && formula < 24.9 {
		fmt.Print("peso normal")
	} else if formula >= 25 && formula < 30 {
		fmt.Print("sobrepeso")
	} else {
		fmt.Print("obesidad")
	}

	fmt.Println("")
	return formula

}
