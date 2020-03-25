package main

import (
	"fmt"
)

func main() {
	//CANALES: permite la comunicacion entre gorutinas

	bodegaOrigen := []string{"Php", "Javascript", "Html", "css", "Java", "Base de Datos"}
	bodegaDestino := []string{}

	//Declaramos 2 canales en el hilo main
	fotocopiadora := make(chan string)
	fotocopiado := make(chan string)

	//Gorutina que se encarga de enviar los libros al canal fotocopiadora
	go func() {
		for _, libro := range bodegaOrigen {
			fmt.Println("1.-Entregado libro a la gorutina a traves del canal fotocopiadora")
			//time.Sleep(time.Second * 2)
			fotocopiadora <- libro
			fmt.Println("2.-Finalizando la entrega de libro a la gorutina a traves del canal fotocopiadora")
		}
	}()

	//Gorutina que se encarga de recibir los libros del canal fotocopiadora y enviar los libros al canal fotocopiado
	go func() {
		for {
			//Entrega el contenido de la fotocopiadora a la variable libro
			fmt.Println("3.-Recibiendo libro del canal fotocopiadora")
			libro := <-fotocopiadora
			fmt.Println("4.-Recibio libro del canal fotocopiadora")

			fmt.Println("5.-Enviando libro al canal fotocopiado")
			//time.Sleep(time.Second * 3)
			fotocopiado <- libro
			fmt.Println("6.-Finalizo envio del libro al canal fotocopiado")

			//Agregar libro a la bodega destino
			bodegaDestino = append(bodegaDestino, libro)
			if len(bodegaDestino) == len(bodegaOrigen) {
				//cerra el canal fotocopiado

				close(fotocopiado)
				//close(fotocopiadora)

			}
		}
	}()

	fmt.Println("Listados de libros fotocopiados")
	//Instrucciones del Hilo main que esta a la espera de los libros a traves de la escucha del canal fotocopiado
	for {
		fmt.Println("8.-a la espera del libro del canal fotocopiado....")
		if libro, ok := <-fotocopiado; ok { // ok es un bool que dice si el canal esta abierto o cerrado
			fmt.Println("9.-Recibiendo libro del canal fotocopiado")
			fmt.Println(libro)
			fmt.Println("10.-Recibio el libro del canal fotocopiado")
		} else {
			break
		}
	}

}
