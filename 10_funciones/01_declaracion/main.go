package main

import "fmt"

func main() {
	saludar("Dickson", 31)
}

func saludar(nombre string, edad uint8) {
	fmt.Println("Hola ", nombre)
	fmt.Printf("Tienes %d años\n", edad)
	if edad > 30 {
		return
	}
	fmt.Println("Eres joven")
}
