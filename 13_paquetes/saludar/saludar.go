package saludar

import "fmt"

// MeVes variable para utilizar fuera del paquete
var MeVes string

// Saludar saluda a una persona
func Saludar(nombre string) {
	fmt.Println("Hola", nombre)
}
