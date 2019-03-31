package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct {
	message string
}

// implementar una interfas del metodo
func (m messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, m.message)
}

func main() {
	// handlers crea los manejadores para las rutas que manejemos
	// es una interfas que contiene un metodo que permite recibir una interfas y un request
	mux := http.NewServeMux()
	m1 := &messageHandler{message: "<h1>Hola desde un handler</h1>"}
	m2 := &messageHandler{message: "<h1>Estos son Gatosr</h1>"}
	m3 := &messageHandler{message: "<h1>Estos son perros</h1>"}

	mux.Handle("/saludar", m1)
	mux.Handle("/gatos", m2)
	mux.Handle("/perros", m3)

	log.Println("Ejucutado server en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))

}
