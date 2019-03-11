package main

import "fmt"

func main() {
  //primer manera de definir variables
  var nombre, apellido string
  nombre= "Dickson"
  apellido = "Garcia"
  //segunda manera de definir variables
  apellido2 := "Rincon"
  //tercera manera de definir variables
  usuario1 , usuario2 := "emely", "joaquin"

  fmt.Println(nombre, apellido, apellido2)
  fmt.Println(usuario1, usuario2)
}
