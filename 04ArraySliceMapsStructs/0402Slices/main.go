package main

import "fmt"

func main() {
	/*
	   Slices
	   1.-Apuntador a un array
	   2.-Tamaño (no es fijo)
	   3.-Capacidad: Cada vez que su capacidad es superada esta se dobla. Ejemplo Al asignar el 5to elemento
	   (Gaspar) y al ver que no hay capacidad ya que con Renee los 4 espacios de memorias ya se habia ocupado
	   entonces lo que hace golang es doblar el capacidad y luego recien coloca a Gaspar en el 5to lugar. Si se
	   asigna otro elemento como hay espacio en memoria no hay necesidad de cambiar la capacidad lo unico que
	   hace es asignar.
	*/

	var nombre []string
	fmt.Printf("Su tamño es: %d y su capacidad es: %d \n", len(nombre), cap(nombre))
	nombre = append(nombre, "Pedro")
	fmt.Printf("Su tamño es: %d y su capacidad es: %d \n", len(nombre), cap(nombre))
	nombre = append(nombre, "Juan")
	fmt.Printf("Su tamño es: %d y su capacidad es: %d \n", len(nombre), cap(nombre))
	nombre = append(nombre, "Rosa")
	fmt.Printf("Su tamño es: %d y su capacidad es: %d \n", len(nombre), cap(nombre))
	nombre = append(nombre, "Renee")
	fmt.Printf("Su tamño es: %d y su capacidad es: %d \n", len(nombre), cap(nombre))
	nombre = append(nombre, "Gaspar")
	fmt.Printf("Su tamño es: %d y su capacidad es: %d \n", len(nombre), cap(nombre))

	fmt.Println(nombre)
	fmt.Printf("%T \n", nombre)

	nombre[3] = "Luis"
	fmt.Println(nombre)

	/*
		make: asignar un slice configurando tamaño y capacidad inicial
		make(tipo de dato del slice, tamaño inicial, capacidad inicial(opcional))
	*/
	apellidos := make([]string, 0, 5)
	fmt.Printf("Su tamño es: %d y su capacidad es: %d \n", len(apellidos), cap(apellidos))
	apellidos = append(apellidos, "Palma")
	fmt.Printf("Su tamño es: %d y su capacidad es: %d \n", len(apellidos), cap(apellidos))
	apellidos = append(apellidos, "Mendez")
	fmt.Printf("Su tamño es: %d y su capacidad es: %d \n", len(apellidos), cap(apellidos))
	fmt.Println(apellidos)

	/*
		declarando y asignando valores a un slice
	*/
	apodos := []string{
		"Chato",
		"Botija",
	}
	fmt.Printf("Su tamño es: %d y su capacidad es: %d \n", len(apodos), cap(apodos))
	fmt.Println(apodos)

}
