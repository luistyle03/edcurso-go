package main

import "fmt"

func main() {

	//Ejecuta indefinidamente el ciclo hasta que usemos el break dentro del ciclo para salir
	a := 0
	for {

		fmt.Println("Pedro ", a)
		a++
		if a == 100 {
			break
		}
	}
}
