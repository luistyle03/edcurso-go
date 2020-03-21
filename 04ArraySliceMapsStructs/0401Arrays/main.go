package main

import "fmt"

func main() {
	var nombre [10]string

	nombre[0] = "Pedro"
	nombre[8] = "Juan"
	nombre[9] = "Rosa"

	apellidos := [3]string{"Luis",
		"Palma",
		"Mendez",
	}

	fmt.Println("Imprimiendo la variable nombre: ", nombre)
	fmt.Println("Imprimiendo valores desde el indice 0 hasta el indice 9-1:8  ", nombre[0:9])
	fmt.Println("Imprimiendo el tamaño del arreglo nombre: ", len(nombre))
	fmt.Println("Imprimiendo la variable apellidos: ", apellidos)
	fmt.Printf("Tipo de dato de la variable nombre: %T \n", nombre)
	fmt.Println("Imprimiendo solo el valor del indice 8 del array nombre: ", nombre[8])
}
