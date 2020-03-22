package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			//ya no continua el resto de instruciones del for y se salta a la siguiente iteracion
			continue
		}
		if i == 8 {
			//sale del for
			break
		}
		fmt.Println("Iteracion ", i)
	}

	arreglo := [3][3]int{}
	valor := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			arreglo[i][j] = valor
			valor++

		}
	}
	fmt.Println(arreglo)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(arreglo[i][j])
		}
	}

}
