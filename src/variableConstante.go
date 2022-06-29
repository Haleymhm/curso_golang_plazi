package main

import "fmt"

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1416
	fmt.Println("Uso y declacacion de Constantes")
	fmt.Println("Valor de Pi con tipo de dato: ", pi)
	fmt.Println("Valor de Pi sin tipo de dato: ", pi2)

	base := 12
	var altura int = 12
	var area, areat int
	area = base * altura
	areat = (base * altura) / 2

	fmt.Println("La base es: ", base, "   La altura:", altura)
	fmt.Println("El area del rectangulo es: ", area)
	fmt.Println("El area del triangulo es: ", areat)

}
