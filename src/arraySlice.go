package main

import (
	"fmt"
)

func main() {
	fmt.Println("Manejo Basico de Array y Slice")

	//Los array tienen numero de posiciones fija, no se puede agregar nuevas posiciones
	// solo cambiar los valores
	var array [4]int

	array[0] = 10
	array[2] = 15
	fmt.Println(array)

	fmt.Println("Tamaño del arrey:", len(array), "capacidad maxima:", cap(array))

	fmt.Println("----------")
	slice := []int{2, 3, 8, 6, 7, 1, 9}

	fmt.Println(slice)

	fmt.Println("Tamaño del slice:", len(slice), "capacidad maxima:", cap(slice))

	fmt.Println(slice[0])   //posicion especifica
	fmt.Println(slice[:3])  // valores inferior a la posicion indicada
	fmt.Println(slice[2:4]) //valores desde la posicion indicada hasta - pos final
	fmt.Println(slice[4:])  // valores desde la pos incicada

	// append: metodo que agraga nuevo valor al slice
	slice = append(slice, 23)
	fmt.Println(slice)
	newSlice := []int{42, 13, 28}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
