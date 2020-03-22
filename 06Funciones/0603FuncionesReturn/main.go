package main

import "fmt"

func main() {
	saludar("Pedro", 30)
	fmt.Println("La suma de 10 y 25 es", sumar(10, 25))
	fmt.Println("La suma2 de 10 y 25 es", sumar(10, 25))
}

func saludar(nombre string, edad uint8) {
	fmt.Printf("Hola %s tienes %d años de edad.\n", nombre, edad)
	if edad > 30 {
		//con return sales de la funcion
		return
	}
	fmt.Println("Eres joven")
}

func sumar(a int8, b int8) int8 {
	return a + b
}

//Haciendo lo mismo que la funcion anterior
//a,b int8 solo se puso asi, ya que a y b son del mismo tipo de dato
func sumar2(a, b int8) int8 {
	return a + b
}
