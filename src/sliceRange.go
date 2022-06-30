package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {

	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	//fmt.Println(textReverse)
	if text == textReverse {
		fmt.Println("El valor", text, " ES PALINDROMO")
		//return "ES PALINDROMO"
	} else {
		fmt.Println("El valor", text, "NO ES PALINDROMO")
	}
}
func main() {
	fmt.Println("Hello World")
	valores := []string{"aMa", "hola", "arePera", "limon"}

	for _, valor := range valores {
		v := strings.ToUpper(valor)
		isPalindrome(v)
	}

}
