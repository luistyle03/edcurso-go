package main

import "fmt"

func main() {
	saludar("Pedro", 30)
}

func saludar(nombre string, edad uint8) {
	fmt.Printf("Hola %s tienes %d años de edad.\n", nombre, edad)
}
