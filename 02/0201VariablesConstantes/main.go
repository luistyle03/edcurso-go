package main

import "fmt"

func main() {

	var apellido2 string
	apellido2 = "Mendez"

	//declara y asigna el tipo de dato lo determina el dato que lo incializa
	nombre, apellido := "Pedro", "Luis"

	nombre = "Alexys"

	//No se puede usar nuevamente el simbolo de declaracion y asignacion(:=)
	//nombre: = "Pedro"
	//Sin embargo se puede utilizar (:=) siempre y cuando al menos uno de ellos se declara y se asigna
	nombre, edad := "Pedro", 20

	fmt.Println(nombre, apellido, apellido2, edad)
}
