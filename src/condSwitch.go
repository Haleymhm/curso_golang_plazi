package main

import (
	"fmt"
)

func esPar(numero int) bool {
	return numero%2 == 0
}

func main() {
	fmt.Println("** CONDICICIONA SWITCH **")
	num := esPar(2)
	switch num {
	case true:
		fmt.Println("El Numero es par")
	default:
		fmt.Println("El Numero es impar")
	}

	fmt.Println("*****")
	var mes int8 = 0
	switch {
	case mes > 0 && mes <= 3:
		fmt.Println("Primer Trimestre")
	case mes > 3 && mes <= 6:
		fmt.Println("Segundo Trimestre")
	case mes > 6 && mes <= 9:
		fmt.Println("Tercer Trimestre")
	case mes > 9 && mes <= 12:
		fmt.Println("Cuarto Trimestre")
	default:
		fmt.Println("Este no es un mes valido")
	}
}
