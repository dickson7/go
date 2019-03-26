package main

import "github.com/dickson7/go/15_interfaces/animales"

func main() {
	var p animales.Perro
	var g animales.Gato

	p.Nombre = "Terry"
	g.Nombre = "firulais"

	//	AdoptarPerro(p)
	//	AdoptarGato(g)
	AdoptarMascota(p)
	AdoptarMascota(g)
}

func AdoptarMascota(m animales.Mascota) {
	m.Comunicarse()
}

/*
func AdoptarPerro(p animales.Perro) {
	p.Comunicarse()
}

func AdoptarGato(g animales.Gato) {
	g.Comunicarse()
}
*/
