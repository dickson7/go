package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// definicion de cuantos procesadores utilizar
	runtime.GOMAXPROCS(1)
	//variable para manejar las gorutinas
	var wg sync.WaitGroup

	numbers := []uint32{134343, 3423, 32423, 24524, 455768, 987, 98765, 6564}
	//se cargan la cantidad de gorutinas
	wg.Add(len(numbers))

	fmt.Println("Comienza el proceso...")

	for _, v := range numbers {
		go func(a uint32) {
			//se restan gorutinas
			defer wg.Done()
			fmt.Println(a, EsPrimo(a))
		}(v)
	}
	//espera que se terminen las gorutinas
	wg.Wait()
	fmt.Println("Termino el proceso")
}

// EsPrimo valida si es primo
func EsPrimo(a uint32) bool {
	c := 0
	var i uint32
	for i = 1; i <= a; i++ {
		if a%i == 0 {
			c++
		}
	}
	if c == 2 {
		return true
	}
	return false
}
