package main

import (
	"fmt"

	depedida "github.com/dickson7/go/13_paquetes/despedida"
	"github.com/dickson7/go/13_paquetes/saludar"
)

func main() {

	saludar.Saludar("Dickson")
	saludar.MeVes = "Esto es un string asignado desde el main a la variable en el paquete"
	fmt.Println(saludar.MeVes)

	nombre := "arley"
	depedida.Despedirse(nombre)
}
