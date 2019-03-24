package main

import "fmt"

func main() {
	/*
		anonima := func() {
			fmt.Println("Me imprimo en una variable")
		}
		anonima()
	*/

	secuencia1 := secuencia()
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println("__________")
	secuencia2 := secuencia()
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
}

func secuencia() func() int32 {
	var x int32
	return func() int32 {
		x++
		return x
	}
}
