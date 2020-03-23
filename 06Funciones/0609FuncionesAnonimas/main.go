package main

import "fmt"

/*
Las funciones anonimas son como un tipo de datos
Puede asignarse a una variable
Se autojecuta
*/
func main() {
	func() {
		fmt.Println("Funcion anonima:Me autoejecuto")
	}()

	anonima := func() {
		fmt.Println("Funcion anonima:me asigno a una variable")
	}
	anonima()

	secuencia1 := secuencia()
	/*secuencia1 es una variable que tiene como valor una funcion, por lo
	que cada vez que llame a la variable secuencia1 su valor que es una funcion sera llamado sin alterar
	el contenido de la funcion (incluye variables). Osea si una varible dentro de dicha funcion vale 3
	esta cuando sea llamada otra vez seguira siendo 3.
	¿PORQUE? Explicacion:
	Por ejemplo si secuencia1 fuera de un tipo de dato int32 cuyo valor es 5 cada vez que se llame
	a dicha variable el valor seguiria siendo 5 (mientras la variable siga viva dentro del scope y no se
	cambie su valor esta sera 5). Analogamente pasa lo mismo cuando	el tipo de dato es una funcion (anonima)
	si en una llamada a secuencia1 el valor de x es 1, cuando se llame a otra vez a secuencia1 el valor
	(que es un funcion)no cambiara ya que aun sigue viva la variable, sin embargo el valor de x aumentara
	uno mas osea sera 2, ¿Porque? porque justamente la funcion sigue viva y por lo tanto el contador x
	sigue aumentando hasta que muera la funcion (osea hasta que muera secuencia 1)

	*/
	fmt.Println(secuencia1()) //x vale 1
	fmt.Println(secuencia1()) //x vale 2
	fmt.Println(secuencia1()) //x vale 3
	fmt.Println("---------")

	secuencia2 := secuencia()
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())

}

//Funcion que retorna una funcion anonima
func secuencia() func() int32 {
	var x int32           //el valor de x no muere entre llamadas
	return func() int32 { //como retorna una funcion anonima, el return debe retornar una funcion anonima
		x++
		return x
	}
}
