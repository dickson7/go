package main

import "fmt"

type Persona struct {
	nombre string
	edad   int8
}

type Numero int

// Firma del metodp
func (p Persona) saludar() {
	fmt.Println("Hola soy una persona")
}

func (n Numero) presentarse() {
	fmt.Println("Soy un numero")
}

func (p *Persona) asignarEdad(i int8) {
	if i < 0 {
		fmt.Println("La edad negativa no es permitida")
	} else {
		p.edad = i
	}
}

func main() {
	var persona Persona
	var numero Numero

	persona.saludar()
	numero.presentarse()

	persona.asignarEdad(-30)
	fmt.Println(persona)
}
