package main

import "fmt"

// or ||
// and &&
// not !

func main() {
	if 5 < 10 {
		fmt.Println("Esto esta en el bloque verdadero")
	}
	//valida la antrada a bloque de codigo
	if isValid := true; isValid {
		fmt.Println("Esto esta entrando en el bloque verdadero")
	} else {
		fmt.Println("Bloque falso")
	}
	fmt.Println("Fin del programa")
}
