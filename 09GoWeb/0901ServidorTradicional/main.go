package main

import (
	"log"
	"net/http"
)

func main() {

	/*Manera no eficiente:
	http.Handle("/", http.FileServer(http.Dir("public")))
	log.Println("Ejecutando servidor en http://localhost:8080...")
	log.Println(http.ListenAndServe(":8080", nil)) //El primer parametro es el ip publico con el que sale a internet seguido por : y luego el puerto. El segundo parametro es el multiplexor que va a servir. }
	//Si el puerto esta ocupado el log mostrara el error en la consola
	*/

	//Servidor de archivos estaticos
	mux := http.NewServeMux()
	//Multiplexor:permite recibir muchas mas solicitudes y procesarlas de manera mas eficiente que la anterior forma, es un ruteador. Go trae uno por defecto que es Mux
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs) //manejador
	log.Println("Ejecutando servidor en http://localhost:8080...")
	log.Println(http.ListenAndServe(":8080", mux))

}
