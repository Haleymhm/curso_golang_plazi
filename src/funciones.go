package main

import (
	"fmt"
)

func msgSimple() {
	fmt.Println("Hello World")
}
func msgSimpleInput(name string) {
	fmt.Println("Hello", name)
}

func msgSimpleInputs(name string, age int) {
	fmt.Println("Hello", name, " y tu edad es:", age)
}

func calcArea(a, b int) int {
	//
	return a * b
}
func returnValues(a, b int, c string) (e, d int) {
	return b, a * 2
}
func main() {
	msgSimple()
	msgSimpleInput("Haleym")
	msgSimpleInputs("Halam", 16)
	area := calcArea(3, 4)
	fmt.Println("El area es:", area)

	n1, _ := returnValues(12, 41, "JJ")
	fmt.Println("Valor devuelto:", n1)
}
