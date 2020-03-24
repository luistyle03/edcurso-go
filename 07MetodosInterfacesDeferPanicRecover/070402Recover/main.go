package main

import "fmt"

func main() {
	f()
}

/*
SI UNA FUNCION ENTRA EN PANICO, UNA VEZ QUE ENTRA EN PANICO NO EJECUTA LA SIGUIENTES INSTRUCCIONES DEL
PROGRAMA Y PROCEDE A EJECUTAR TODOS LOS DEFER QUE ESTUVIERON ANTES DEL PANIC. LOS DEFER DESPUES
DEL PANIC NO SE EJECUTA.

SI HAY UN PANIC Y TENEMOS UNA INSTRUCCION DE RECOVER DENTRO DE DEFER (OjO: linea de defer esta antes que
panic), ENTONCES LA LINEA DEL PANIC NO SE EJECUTA Y SE VA DE FRENTE A RECOVER CUYO VALOR DEVUELTO
ES EL MISMO TEXTO o NUMERO QUE PUSIMOS EN PANIC(texto o numero)

*/

func f() {
	defer func() { //este defer tiene que ir antes de lo que posiblemente fallara, si lo  hubierAmos
		//puesto este defer al final de la funcion f nunca entraria ya que el panic estaria antes.
		//Recover se usa cuando por ejemplo falla una conexion a base de datos entra en panico
		//el recover toma el control y podria probar la conexion en una segunda base de datos para
		//que el programa continue
		if r := recover(); r != nil { //Cuando ocurre panic recover entra en accion. Recover tiene
			//el mensaje de panic
			fmt.Printf("%T\n", r)
			fmt.Printf("El panico con mensaje: 'Nooooo!!!. No puede ser mayor a %d'. se recupero a traves de recover \n", r)
		}
	}()
	fmt.Println("Llamando g.")
	g(5)
	fmt.Println("Retorna normalmente desde g")

}

func g(i int) {
	if i > 3 {
		fmt.Println("AAAAAaaaaa")
		panic(i)
	}

	defer fmt.Println("Defer en la funcion g")
	fmt.Println("Imprimiendo en g", i)
}
