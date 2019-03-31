package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// con gorm se puede pasar una estructura con los campos de la estructura
// esa estructura debe tener unos campos especificos (gorm.Model)

// Producto contiene los datos de un articulo
type Producto struct {
	gorm.Model   //ID, CreateAt, UpdatedAt, DeleteAt
	CodigoBarras string
	Precio       uint8
}

func main() {
	db, err := gorm.Open("mysql", "root:@/basedatosprueba?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Error en la conexión a la base de datos")
	}
	defer db.Close()
	fmt.Println("Se conectó a la base de datos")

	// creamos una tabla pasadole una estructura en un puntero
	/*
		db.CreateTable(&Producto{})
		fmt.Println("Se creo la tabla productos en la BD")
	*/

	// pasar datos a una tabla desde la estructura
	/*
		db.Create(&Producto{
			CodigoBarras: "022139043944",
			Precio:       200,
		})
		fmt.Println("Se creo registro en base de datos")
	*/

	// consultar datos de la base de datos
	// trae el primer registro
	/*
		var p Producto
		db.First(&p)
		fmt.Println(p)
		fmt.Println("Precio del producto consultado:", p.Precio)
	*/

	// cuando hay mas de un registor y queremos consultar no el 1ro sino el 2do registro
	/*
		var p Producto
		// consulta el registro con el id numero 2
		db.First(&p, 2)
		fmt.Println(p)
		fmt.Println("Precio del producto consultado:", p.Precio)
	*/

	// consultar un registro y modificar su Precio
	var p Producto
	db.First(&p, 2)
	fmt.Println(p)
	fmt.Println("Precio del producto consultado:", p.Precio)

	p.Precio = 50
	db.Save(&p)
	fmt.Println("El nuevo precio de producto es:", p.Precio)
}
