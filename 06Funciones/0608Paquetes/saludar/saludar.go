//la carpeta y package deben tener el mismo nombre
package saludar

import "fmt"

var MeVes string
var noMeVes string

/*Toda funcion a exportar(osea funcion que puede ser llamada desde otros paquetes),
debe tener la primera letra en mayuscula
*/
func Saludar(nombre string) {
	fmt.Println("Hola", nombre)
}
