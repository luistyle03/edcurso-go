package main

import "fmt"

func main() {
	a := 3

	/*Puntero: es la direccion de la memoria de una variable. Para conocer la direccion
	ponemos &variable . Ejemplo:0xc000014108
	Un puntero vacio contiene NIL
	Si le ponemos el * antes de una direccion de un puntero esta nos devuelve su valor.
	Ejemplo: *puntero es igual al valor	donde la apunta la direccion del puntero
	*/
	fmt.Println(&a)
	fmt.Println(*&a)

	fmt.Println("PASE POR VALOR: copia el valor de una variable a otra")
	b := 3
	fmt.Println("Antes de duplicar, b=", b)
	duplicar(b)
	fmt.Println("Despues de duplicar, b=", b)

	fmt.Println("PASE POR REFERENCIA: Usamos punteros para ello se copia la direccion de la variable a otra variable, pero al momento de duplicar, esta dobla el valor a donde apunta la direccion del puntero")
	c := 3
	fmt.Println("Antes de duplicar, valor de c=", c)
	duplicarporpuntero(&c)
	fmt.Println("Despues de duplicar, valor c=", c, "Se llego a duplicar debido a que en todo momento se trabajo con el mismo valor de la variable, todo lo contrario con el pase con valor ya que la funcion copio en una nueva variable por lo que cualquier cambio que sucedio dentro de la funcion, se quedo ahi")

}

func duplicar(x int) {
	x = x * 2
	fmt.Println("Dentro de la funcion duplicar, x vale ", x)
}

func duplicarporpuntero(x *int) {

	*x *= 2 //multiplica el valor de x por 2
	//Seria otra forma:(Recuerda que *puntero accede al valor donde apunta la direccion de dicho puntero)
	//*x = *x * 2
	fmt.Println("Dentro de la funcion duplicar,puntero x vale ", *x)
}
