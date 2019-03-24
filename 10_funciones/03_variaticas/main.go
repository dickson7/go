package main

import "fmt"

func main() {
	saludarVarios(10, "dickson", "emely", "jaoquin")
}

//las funcions variaticas solo recibe un unico parametro variatico
func saludarVarios(edad uint8, nombres ...string) {
	// para saber que tivo de datos es una variable
	fmt.Printf("%T\n", nombres)
	for _, v := range nombres {
		fmt.Println("Hola", v, "edad", edad)
	}
}
