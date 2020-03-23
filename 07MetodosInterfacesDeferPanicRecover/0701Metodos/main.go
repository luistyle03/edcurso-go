package main

import "fmt"

type Persona struct {
	nombre string
	edad   int8
}

//Asignando metodo saludar al struct Persona. Seria como el metodo de una clase
func (p Persona) saludar() {
	fmt.Println("Hola soy una persona")
}
func (p *Persona) asignarEdad(e int8) {
	if e >= 0 {
		p.edad = e
	} else {
		fmt.Println("No es una edad valida. Debe ser mayor o igual a cero")
	}
}

type Numero int64

//Asinando metodo a un tipo de dato entero64
func (n Numero) presentarse() {
	fmt.Println("Hola soy un numero")
}

func main() {

	//Metodos de un struct
	var persona Persona
	persona.saludar()
	persona.asignarEdad(30)
	//persona.asignarEdad(-35) //No le asignaria la edad ya que el metodo esta validado
	fmt.Println(persona)

	//Metodos de un entero
	var numero Numero
	numero.presentarse()
}
