package main

import "fmt"

func main() {
	bodegaOrigen := []string{"php", "javascript", "html", "css", "java", "BD", "git"}
	bodegaDestino := []string{}

	fotocopiadora := make(chan string)
	fotocopiado := make(chan string)

	go func() {
		for _, libro := range bodegaOrigen {
			fotocopiadora <- libro
		}
	}()

	// se carga fotocopiadora
	go func() {
		for {
			// se entrega el contenido de fotocopiadora a var libro
			libro := <-fotocopiadora
			// se pasa el libro a fotocopiado
			fotocopiado <- libro
			// agregar el libro a la bodega destino
			bodegaDestino = append(bodegaDestino, libro)

			if len(bodegaOrigen) == len(bodegaDestino) {
				// cerrar el canal fotocopiado
				close(fotocopiado)
			}

		}
	}()

	fmt.Println("Listado de libros fotocopiados")
	for {
		if libro, ok := <-fotocopiado; ok {
			fmt.Println(libro)
		} else {
			break
		}
	}

}
