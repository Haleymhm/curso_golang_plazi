package main

import "fmt"

func main() {
	fmt.Println("** KEYWORDS **")
	defer fmt.Println("** FIN DEL KEYWORDS **")
	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 4 {
			fmt.Println("Soy el 4")
			fmt.Println("despues de continue")
			continue
		}

		// Break
		if i == 8 {
			fmt.Println("Se detiene el ciclo con un break antes de llagar a 10")
			break
		}
	}
}
