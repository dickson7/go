package main

import "fmt"

func main() {
	f()
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%T\n", r)
			fmt.Println("Recuperado en f:", r)
			fmt.Println("Se llama a la funcion de respaldo para continuar la app")
		}
	}()

	fmt.Println("Llamando a g.")
	g(5)
	fmt.Println("Retorna normalmente desde g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("paaaaaaaaanicooooo")
		panic("el numero no puede ser mayor a 3")
	}
	defer fmt.Println("Defer en la funcion g")
	fmt.Println("Imprimiendo en g:", i)
}
