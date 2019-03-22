package main

import "fmt"

func main() {
	// for infinito se ejecuta infinitamente y no hay
	//  ninguna condicion que lo detenga
	// Se utiliza para escuchar permanentemente conexiones
	// a la base de datos, sockets etc
	for {
		fmt.Println("Esto no se va a detener")
	}
}
