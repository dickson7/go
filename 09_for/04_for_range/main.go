package main

import "fmt"

func main() {

	numeros := []string{"cero", "uno", "dos", "tres"}

	// el range devuelve dos valores el indice y el valor
	for i, v := range numeros {
		fmt.Println(i, v)
	}

	//si no queremeos utilizar el indice de coloca _
	for _, v := range numeros {
		fmt.Println(v)
	}

	// El valor se puede omitir en un range
	for i := range numeros {
		fmt.Println(i)
	}

	//ejemplo con mapas funciona de la misma forma
	nombres := map[string]string{
		"es": "España",
		"co": "Colombia",
		"br": "Brasil",
	}
	for i, v := range nombres {
		fmt.Println(i, v)
	}

	// Iteracion de string
	frase := "Hola Mundo"
	for posicion, letra := range frase {
		//fmt.Println(posicion, letra)
		fmt.Println(posicion, string(letra))
	}

	// Iterar slices de entreros
	for _, entero := range []int{15, 36, 24, 87} {
		fmt.Println(entero)
	}

	// Iterar un string
	for _, letra := range "Hola Mundo" {
		fmt.Println(string(letra))
	}
}
