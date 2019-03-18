package main

import "fmt"

func main() {
	//slices
	//tiene un apuntador a un arrays
	// tienen un tamaño
	//tambien tiene una capacidad

	//Declaracion de slices
	//	var nombres []string

	//otra manera de declarar slices
	nombres := make([]string, 0)
	fmt.Printf("Su tamaño es: %d y scapacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Dickson")
	fmt.Printf("Su tamaño es: %d y scapacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Arley")
	fmt.Printf("Su tamaño es: %d y scapacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Emely")
	fmt.Printf("Su tamaño es: %d y scapacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Yoxandra")
	fmt.Printf("Su tamaño es: %d 1y scapacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Joaquin")
	fmt.Printf("Su tamaño es: %d 1y scapacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Gerald")

	//Otra forma de agregar informacion a nuestro slices
	apellidos := []string{
		"Garcia",
		"Zerpa",
		"Rincon",
	}

	fmt.Println(nombres)
	fmt.Println(apellidos)
}
