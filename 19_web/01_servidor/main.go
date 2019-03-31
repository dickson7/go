package main

import (
	"log"
	"net/http"
)

func main() {

	// servidor web de forma tradicional
	/*
		// Ruta a nuestro servidor que va a ser la raiz "/" y que debe hacer al realizar esa petición
		http.Handle("/", http.FileServer(http.Dir("public")))
		log.Println("Ejucutado server en http://localhost:8080")
		// se carga el servidor diciendole que se sirva en un puerto
		// si el puerto esta ocupado va adevolver un error y para capturar ese error lo hacemos con un log
		log.Println(http.ListenAndServe(":8080", nil))
	*/

	// servidor web mejorado con un multiplexor(ayuda a resolver las solucitudes)
	// server mux resuelve mas peticiones
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	log.Println("Ejucutado server en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
