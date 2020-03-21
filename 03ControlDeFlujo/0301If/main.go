package main

import "fmt"

func main() {
	/*
		&& and
		|| or
		! Not
		!= diferente
		== igual

	*/

	//isValid solo tiene alcance dentro de la condicional
	if isValid := true; !(5 > 10) && 1 > 3 && isValid {
		fmt.Println("Es verdadero")
		if 3 < 5 {
			fmt.Println("Es verdadero la condicional interna")
		} else {
			fmt.Println("Es falsa la condicional intena")
		}
	} else if isValid && 8 < 3 {
		fmt.Println("Es falso")
	} else {
		fmt.Println("Recontra falso")
	}

	//isValid no es valido aqui por tema de scope
	//fmt.Println(isValid)

}
