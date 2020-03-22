package main

import "fmt"

func main() {

	//ITERANDO SLICES
	numeros := []string{"uno", "dos", "tres"}
	//i almacen el indice y v el valor
	for i, v := range numeros {
		fmt.Printf("Indice %d Valor %s \n", i, v)
	}

	//Si no deseo almacenar el indice entonces en su posicion ponemos el underline _
	for _, v := range numeros {
		fmt.Printf("Valor %s \n", v)
	}

	//Si no deseo almacenar el valor entonces simplemente no ponemos ,v
	for i := range numeros {
		fmt.Printf("Indice %d \n", i)
	}

	//ITERANDO MAPS
	nombres := map[string]string{"es": "españa", "fr": "Francia"}
	for i, v := range nombres {
		fmt.Printf("Indice %s Valor %s \n", i, v)
	}

	//ITERANDO CARACTERES DE UN STRING
	frase := "Hola Mundo"
	for i, v := range frase {
		//v almacena el ascii
		fmt.Printf("Indice %d Valor %d \n", i, v)
		//string(ascii): nos da el caracter
		fmt.Printf("Indice %d Valor %s \n", i, string(v))
	}

	//ITERANDO UN SLICE CREADO EN EL MISMO FOR
	for i, v := range []int{1, 34, 454, 2, 45} {
		fmt.Printf("Indice %d Valor %d \n", i, v)

	}

	//ITERANDO UN STRING CREADO EN EL MISMO FOR
	for i, v := range "Hola Mundo" {
		fmt.Printf("Indice %d Valor %s \n", i, string(v))

	}
}
