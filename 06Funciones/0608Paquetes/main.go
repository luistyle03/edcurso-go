package main

import (
	"fmt"

	"github.com/luistyle03/edcurso-go/06Funciones/0608Paquetes/despedir"
	"github.com/luistyle03/edcurso-go/06Funciones/0608Paquetes/saludar"
)

func main() {
	nombre := "Pedro"
	saludar.Saludar(nombre)
	saludar.MeVes = "esto es un texto"
	fmt.Println(saludar.MeVes)

	despedir.Despedirse(nombre)
}
