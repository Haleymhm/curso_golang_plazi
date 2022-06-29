package main

import (
	"fmt"
)

func main() {
	// Ciclo for Regular
	for i := 0; i <= 10; i++ {
		fmt.Println("Hello World", i)
	}

	//Ciclo While
	fmt.Println("---")
	cont := 0
	for cont <= 5 {
		fmt.Println("Hello World", cont)
		cont++
	}
	fmt.Println("---")
	for j := 4; j >= 0; j-- {
		fmt.Println("Hello World", j)
	}
}
