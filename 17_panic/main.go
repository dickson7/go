package main

import "fmt"

func main() {
	division(3, 0)
}

func division(dividendo, divisor int) {
	defer fmt.Println("Esto se ejecutara pase lo que pase")
	if divisor == 0 {
		panic("No se puede dividir por cero")
	}
	fmt.Println(dividendo / divisor)

}
