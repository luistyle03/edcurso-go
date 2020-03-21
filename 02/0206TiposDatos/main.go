package main

import "fmt"

func main() {
	//si solo se pone int, de acuerdo a nuestra arquitectura en nuestro caso sera 64
	var a int
	var b int8

	a = 123456
	b = 5

	//Utilizo casting de int8 a int
	c := a + int(b)
	fmt.Println(c)

	d := "Pedro"
	//%s imprime string y %d imprime numeros
	fmt.Printf("Hola %s tienes %d de edad \n", d, c)
	//%T imprime el tipo de dato
	fmt.Printf("C tiene un tipo de datos %T \n", c)

	var nombre string //por defecto string tiene ""
	var numero int    //por defecto int tiene 0
	var foco bool     //por defecto bool tiene false

	fmt.Println(nombre, numero, foco)
}
