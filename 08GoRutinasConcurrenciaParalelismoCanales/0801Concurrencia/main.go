package main

import (
	"fmt"
	"time"
)

func main() {
	var h string
	go MostraNumeros() //con go ejecutamos una gorutina(aplica concurrencia), es decir esta
	//actividad lo hace independientemente en un proceso distinto al main. Cabe mencionar
	//que si se acaba de ejecutar todas las instrucciones del main dicho proceso independiente
	//de la gorutina se terminaria tambien.
	//En nuestro ejemplo si digitamos un texto y el programa ejecuta hasta la ultima instrucion
	// del main (fmt.Println("Digitaste", h) entonces termina el main y termina la gorutina
	// asi no haya terminado de imprimir los 10 numeros que se le encomendaron

	fmt.Println("Digita=> ")
	fmt.Scan(&h)
	fmt.Println("Digitaste", h)
}

func MostraNumeros() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second) //espera un segundo antes de continuar con las siguientes instrucciones
	}
}
