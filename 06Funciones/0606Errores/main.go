package main

import (
	"errors"
	"fmt"
)

func main() {
	//los errores son como un tipo de dato
	//El valor por defecto de un error es nil (nulo)

	err1 := errors.New("Mi mensaje de error")
	fmt.Printf("Tipo de dato de error es %T\n", err1)
	fmt.Printf("El error es %s\n", err1)
	fmt.Println()

	r, err := division(100, 2)
	//fmt.Printf("El error de dividir  es %s\n", err)
	if err != nil { //si hubo error esta variable NO esta vacia (No contiene nil o NULO)
		fmt.Println("Error", err)
		return
	}
	fmt.Printf("El resultado de dividir es %f\n", r)

}

func division(dividendo, divisor float64) (resultado float64, err error) {
	if divisor == 0 {
		err = errors.New("No se puede dividir entre cero")
	} else {
		resultado = dividendo / divisor
	}
	return
}
