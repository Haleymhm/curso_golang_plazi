package main

import (
	"fmt"
)

type car struct {
	brand string
	year  int
	model string
	color string
}

func main() {
	fmt.Println("MAnejo de Estructuras")
	var cars car

	cars.brand = "Chery"
	cars.year = 2022
	cars.model = "Tiggo 3"
	cars.color = "Negro"

	fmt.Println(cars)
}
