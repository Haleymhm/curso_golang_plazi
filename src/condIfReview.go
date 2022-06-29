package main

import (
	"fmt"
)

func esPar(numero int) bool {
	return numero%2 == 0
}
func isValidUser(userName, pass string) bool {
	return userName == "Jose" && pass == "J1234"
}

func main() {
	fmt.Println("Reto de la Clase de condiciona IF")
	if esPar(7) {
		fmt.Println("El Numero es PAR")
	} else {
		fmt.Println("El Numero es IMPAR")
	}

	if isValidUser("Jose", "J12345") {
		fmt.Println("Datos Correctos")
	} else {
		fmt.Println("Datos Incorrectos")
	}
}
