package main

import "fmt"

func main() {
	fmt.Println("** KEYWORDS **")
	defer fmt.Println("** FIN DEL KEYWORDS **")
	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 4 {
			fmt.Println("Es 4")
			continue
		}

		// Break
		if i == 4 {
			fmt.Println("Break")
			break
		}
	}
}
