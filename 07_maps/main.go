package main

import "fmt"

func main() {
	//Declaracion de maps
	/*
		idiomas := make(map[string]string)
		idiomas["es"] = "Español"
		idiomas["en"] = "Inglés"
		idiomas["fr"] = "Frecés"
	*/

	//otra forma de Declaracion de maps
	idiomas := map[string]string{
		"es": "Español",
		"en": "Inglés",
		"fr": "Francés",
	}
	//imprimir todo el mapa
	fmt.Println(idiomas)
	//imprimir solo una posicion en especifico
	fmt.Println(idiomas["fr"])
	//eliminar una posicion del mapa
	delete(idiomas, "en")
	fmt.Println(idiomas)

	// comprobar si una posicion existe
	if idioma, ok := idiomas["en"]; ok {
		fmt.Println("La posicion en si existe y vale ", idioma)
	} else {
		fmt.Println("La posicion NO existe")
	}

	//declaracion de maps con int y string
	numeros := map[int]string{
		1:  "un numero pequeño",
		20: "otro numero",
		-3: "llave negativa",
	}
	fmt.Println(numeros)
}
