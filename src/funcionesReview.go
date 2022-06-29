package main

import (
	"fmt"
)

func calAreaRec(a, b int) int {
	return a * b
}
func calAreaCuadrada(a int) int {
	return a * a
}
func calAreaTriangulo(a, b float32) float32 {
	return (a * b) / 2
}

func calAreaCirculo(a float32) float32 {
	const pi float32 = 3.1416
	return pi * (a * a)
}
func main() {
	areaRec := calAreaRec(3, 7)
	fmt.Println("El area del Rectancgulo es:", areaRec)

	areaCuadrado := calAreaCuadrada(4)
	fmt.Println("El area del Cuadrado es:", areaCuadrado)

	areaTriangulo := calAreaTriangulo(3, 7)
	fmt.Println("El area del Triancgulo es:", areaTriangulo)

	areaCirculo := calAreaCirculo(4)
	fmt.Println("El area del Circulo es:", areaCirculo)

}
