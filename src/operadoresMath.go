package main

import "fmt"

func main() {
	// Declaracion de constantes
	const pi float64 = 3.1416
	var radio float64 = 5.0
	var base int = 8
	var baseM int = 14
	var altura int = 12

	fmt.Println("Operadores Matematicos Basicos")
	fmt.Println("Base:", base, "Altura", altura)
	resultado := base + altura
	fmt.Println("Resul Suma:", resultado)

	resultado = base - altura
	fmt.Println("Resul Resta:", resultado)

	resultado = base * altura
	fmt.Println("Resul Multipicacion:", resultado)

	resultado = base / altura
	fmt.Println("Resul Div:", resultado)

	resultado = base % altura
	fmt.Println("Resul Resto:", resultado)

	fmt.Println("------------------------------")
	fmt.Println("CALCULO DE AREAS")
	fmt.Println("Base:", base, "Altura", altura, "PI:", pi)
	area := base * altura
	fmt.Println("Area de un rectangulo:", area)

	area = (base * altura) / 2
	fmt.Println("Area de un triangulo:", area)

	areaC := pi * (radio * radio)
	fmt.Println("Area de un circulo:", areaC)

	fmt.Println("Base Menor:", base, "Base Mayor", baseM, "Altura:", altura)
	areaT := ((base + baseM) * altura) / 2
	fmt.Println("Area de un traperio:", areaT)
}
