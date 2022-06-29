package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	v1 := 3
	v2 := 2
	fmt.Println("** Condicionales **")
	if v1 > v2 {
		fmt.Println("VALOR 1 ES MAYOR")
	} else {
		fmt.Println("Valor 2 es MAYOR")
	}

	valor, err := strconv.Atoi("25")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("El Valor es:", valor)

	valor2, err2 := strconv.Atoi("ojo")
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("El Valor es:", valor2)
}
