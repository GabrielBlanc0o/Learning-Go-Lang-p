package main

import "fmt"

func main() {
	ejercicio1()
	fmt.Print("\n")
	ejercicio2()
	fmt.Print("\n")
	ejercicio3()
	ejercicio4()
	ejercicio5()
	ejercicio6()
	ejercicio7()
	ejercicio8()
}

func ejercicio1() {
	for i := 1; i < 11; i++ {
		fmt.Print(i, "\n")
	}

}

func ejercicio2() {
	for i := 10; i >= 0; i-- {
		fmt.Print(i, "\n")
	}

}

func ejercicio3() {
	suma := 0

	for i := 0; i <= 100; i++ {
		suma += i
	}
	fmt.Print("La suma del 1 al 100 es: ", suma, "\n")
}

func ejercicio4() {
	contador := 1

	for i := 0; contador <= 10; i++ {
		fmt.Println(contador)
		contador++
	}
	fmt.Println("\n")
}

func ejercicio5() {
	x := 1

	for {
		if x > 10 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("\n")

}

func ejercicio6() {

	var numero int
	fmt.Println("Ingresa numero: ")
	fmt.Scanln(&numero)

	for i := 1; i < 10; i++ {
		mult := numero * i
		fmt.Println(numero, " x ", i, " = ", mult)
	}
	fmt.Println("\n")
}

func ejercicio7() {

	for i := 1; i <= 20; i++ {
		i++
		if i%2 == 0 {
			println(i)
		} else {
			println("")
		}
	}

}

func ejercicio8() {

	var numero int
	res := 1
	fmt.Print("Ingresa un numero : ")
	fmt.Scan(&numero)

	for i := 1; i <= numero; i++ {

		res = res * i
	}
	fmt.Print("El factorial de ", numero, " es: ", res)
}
