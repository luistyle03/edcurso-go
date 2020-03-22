package main

import "fmt"

func main() {
	saludarVarios(30, "Pedro", "Juan", "Rosa")
}

//variablevariatica...string variable vendria a ser un slice del tipo de datos asignado, en este caso slice de strings
//La variable variatica solo puede estar al final de los argumentos
func saludarVarios(edad uint8, nombres ...string) {
	fmt.Printf("El tipo de datos de nombres es %T\n", nombres)
	for _, v := range nombres {
		fmt.Printf("Hola %s\n", v)

	}
	fmt.Printf("Todos tienen %d años de edad \n", edad)
}
