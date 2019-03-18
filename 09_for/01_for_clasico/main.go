package main

import "fmt"

func main() {
	// ejemplo 1
	/*
		  count := 5
				for i := 0; i < count; i++ {
					fmt.Println(i)
				}
	*/

	// ejemplo 2
	/*
				for i := 5; i >= 0; i-- {
					if i == 3 {
		      fmt.Println("Se salta el valor 3")
						continue
					}
					fmt.Println(i)
				}
	*/

	// ejemplo 3
	/*
		count := 5
		for i := 0; i < count; i++ {
			if i == 2 {
				fmt.Println("Cuando i vale 2 se rompe el ciclo")
				break
			}
			fmt.Println(i)
		}
	*/

	// ejemplo 4
	matriz := [3][3]int{}
	valor := 0
	for i := 0; i < 3; i++ {
		for y := 0; y < 3; y++ {
			valor++
			matriz[i][y] = valor
		}

	}
	for i := 0; i < 3; i++ {
		for y := 0; y < 3; y++ {
			fmt.Println(matriz[i][y])
		}
	}
}
