package main

import (
	"log"
	"net/http"
	"text/template"
)

type Persona struct {
	Nombre string
	Edad   uint8
}

func renderizar(w http.ResponseWriter, r *http.Request) {
	p := &Persona{"Dickson", 30}
	// el paquete tmplate procesa los archivos devulve el template y un error y se controlan
	t, err := template.ParseFiles("./views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, p)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderizar)

	log.Println("Ejucutado server en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
	log.Println("Servidor no esta corriendo")
}
