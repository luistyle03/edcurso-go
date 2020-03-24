package main

import "fmt"

/*
	Los defer se ejecutan al final de la funcion sin importar en que numero de linea de instrucciones
	se encuentren, vale decir que se ejecutaran recien cuando acabe de ejecutar todas las
	instrucciones o cuando haya un return o cuando entre en panico la funcion.

	En el caso de panico solo ejecutara los defer que estan antes del panico

	los defer siempre se ejecutan en orden inverso: en este caso una vez acabado main se ejecuta primero cerraSetDatos y luego cerrarBaseDatos
*/

func main() {
	fmt.Println("Conectamos a la base de datos..")

	//Falle o no la conexion a la base de datos, una vez que acabe todo la funcion main, cerrarBaseDatos se ejecutara
	defer cerrarBaseDatos()

	fmt.Println("Consultamos informacion, set de datos")
	//Falle o no la conexion la consulta de datos, una vez que acabe toda la funcion main, cerrarsetDatos se ejecutara
	defer cerraSetDatos()

}

func cerrarBaseDatos() {
	fmt.Println("Cerrar la base de datos")

}

func cerraSetDatos() {
	fmt.Println("Cerrar set de datos")
}
