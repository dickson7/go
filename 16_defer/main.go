package main

import "fmt"

func main() {
	// Defer es como una sala de espera
	// que hace la ejecucion de codigo cuando se termina la funcion

	fmt.Println("Conectando a la base de datos...")
	defer cerrarDB()

	fmt.Println("Consultamos informacion de set de datos")
	defer cerrarSetDeDatos()
}

func cerrarDB() {
	fmt.Println("Cerrar la BD")
}

func cerrarSetDeDatos() {
	fmt.Println("Cerrar el set de datos")
}
