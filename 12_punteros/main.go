package main

import "fmt"

func main() {
	a := 3
	fmt.Println("antes de duplicar, a =", a)
	duplicar(&a)
	fmt.Println("Despues de duplicar, a =", a)
}

func duplicar(a *int) {
	*a *= 2
	fmt.Println("Dentro de la funcion deplicar a =", *a)
}
