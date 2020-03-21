package main

import "fmt"

func main() {
	a := 1

	switch a {
	case 1:
		fmt.Println("Es uno")
		//no es necesario usar break como en otros programas, una vez que entra a algun AUTOMATICAMENTE sale del switch
	case 2:
		fmt.Println("Es dos")
	default:
		fmt.Println("Es otro numero")
	}

	//declaro una variable cuyo alcance solo es el switch
	switch b := 30; {
	case b < 18:
		fmt.Println("Es menor de edad")
	case b < 45:
		fmt.Println("Es adulto")
	default:
		fmt.Println("Es de tercera edad")
	}

	c := 5
	switch c {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("No es un dia de la semana")
	}
	switch c {
	case 1:
		fallthrough
		//si entra aqui fallthrough obliga a que evalue el siguiente caso, osea case 2
	case 2:
		fallthrough
		//si entra aqui fallthrough obliga a que evalue el siguiente caso, osea case 3
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Println("Estas entre semana")
		//Imprime pero como no hay fallthrough sace del switch
	case 6:
		fallthrough
	case 7:
		fmt.Println("Estas en fin de semana")
	default:
		fmt.Println("No es un dia de la semana")
	}

}
