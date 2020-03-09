package main

import "fmt"

func main() {

	//Declarando y luego asignado valor a una variables
	var apellido1 string
	apellido1 = "Mendez"

	//Declarando y luego asignado valores a varias variables
	var nombre2, apellido2 string
	nombre2 = "Pedro"
	apellido2 = "Rodriguez"

	//declara y asigna en una sola linea. El tipo de dato lo determinar el valor de inicializacion de la variable
	nombre3, apellido3 := "Pedro", "Luis"

	//No se puede usar nuevamente el simbolo de declaracion y asignacion(:=)
	//nombre3: = "Pedro"

	//Sin embargo se puede utilizar solo para asignar (:=), siempre y cuando al menos en la misma linea uno de ellos se declare y asigne (en este caso edad)
	nombre3, edad3 := "Alexyx", 20

	fmt.Println(apellido1, nombre2, apellido2, nombre3, apellido3, edad3)

}
