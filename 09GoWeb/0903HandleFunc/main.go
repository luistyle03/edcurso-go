package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
HANDLEFUNC son como tipo de datos que convierte funciones en manejadores http, para ello la firma debe
tener 2 parametros responsewriter y request. En este caso no es necesario crear un struct y un metodo,
solo basta con una funcion
*/

//Es como si fuera nuestro controlador
func messageHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hola esto es un handlerfunc</h1>")
}

func messageHFCustom(message string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, message)
		},
	)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", messageHandlerFunc)
	mux.Handle("/saludar", messageHFCustom("<h1>Hola desde Handle Function con mensaje personalizado</h1>"))
	mux.Handle("/despedir", messageHFCustom("<h1>Bye desde Handle Function con mensaje personalizado</h1>"))

	log.Println("Ejecutando servidor en http://localhost:8080...")
	log.Println(http.ListenAndServe(":8080", mux))

}
