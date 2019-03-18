package main

import "fmt"

// Persona es una estructura
type Persona struct {
	Nombre string
	Edad   uint8
	Emails []string
}

func main() {
	// 1. forma de declarar y agregar campos a la estructura
	/*
		var persona1 Persona
		persona1.Nombre = "Dickson"
		persona1.Edad = 30
		fmt.Println(persona1)
		fmt.Println(persona1.Edad)
	*/

	// 2. Manera con shorthand
	/*
		persona2 := Persona{}
		persona2.Nombre = "Dickson"
		persona2.Edad = 30
		fmt.Println(persona2)
	*/

	// 3. forma con los datos dentro de las llaves
	/*
	     persona2 := Persona{
	   		Nombre: "Dickson",
	   		Edad:   31,
	   	}
	   	fmt.Println(persona2)
	*/

	// 4. forma con los datos dentro de las llaves
	persona2 := Persona{
		"Dickson",
		32,
		[]string{"a@b.com", "d@b.com"},
	}
	fmt.Println(persona2)
}
