package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct {
	message string
}

//ServeHTTP: nombre obligatorio que implementara el metodo ServeHTTP
func (m messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //1er parametro respuesta al usuario, 2do parametro solicitud del usuario
	fmt.Fprintf(w, m.message)
}

func main() {

	//Handle o manejador para la rutas de nuestra aplicacion.
	//Handle es una interfaz que contiene un metodo que nos permiter escribir una respuesta o recibir una solicitud del usuario a traves de sus 2 parametros
	mux := http.NewServeMux()
	m1 := &messageHandler{message: "<h1> Hola desde un handler<h1>"} //esto es codigo html que devolvera el server
	m2 := &messageHandler{message: "<h1> Hola te saluda un perro<h1>"}
	m3 := &messageHandler{message: "<h1> Hola te saluda un gato<h1>"}

	mux.Handle("/saludar", m1) //Handle() se le dara como primer parametro la ruta (localhost:8080/SALUDAR) y como segundo paremetro lo que devolvera al usuario
	mux.Handle("/perro", m2)
	mux.Handle("/gato", m3)

	log.Println("Ejecutando servidor en http://localhost:8080...")
	log.Println(http.ListenAndServe(":8080", mux))

}
