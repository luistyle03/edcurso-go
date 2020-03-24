package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	/*
		fmt.Println("Inicio el proceso sin gorutinas")
		//OJO: Aun colocando runtime.GOMAXPROCS(8) siempre trabaja un procesador(Mi pc tiene 8 procesadores pero igual el proceso es realizado por uno solo, sin embargo cabe mencionar que llevar a cabo todo el proceso puede ser realizado por diferentes procesadores, pero siempre solo UNO hace el trabajo)
		for i := 1; i <= 1000000; i++ {
			fmt.Println(i, esPrimo(i))

		}
		fmt.Println("Termino el proceso")
	*/

	fmt.Println("Inicia el proceso APLICANDO CONCURRENCIA(goRutinas) Y PARALELISMO(mas de un procesador trabajando en simultaneo con varias goRutinas). Se hara uso de sync.WaitGroup para que hilo principal main espere a todas las gorutinas lanzadas")
	runtime.GOMAXPROCS(8) //Mi pc tiene 8 procesadores, todas estaran trabajando en simultaneo diferentes gorutinas (ESTO ES PARALELISMO)
	var wg sync.WaitGroup //grupo de espera, es una cola de gorutinas con el cual GO sabra si hay gorutinas en espera o procesando instrucciones. Si la cola aun esta llena el hilo principal main no finalizara el programa ya que habra gorutinas pendientes(Se hara uso de .Done() para quitar de la cola una gorutina que ya termino y .Wait() para decirle que el hilo principal main no puede seguir con las siguientes instrucciones en dicho hilo hasta que se finalicen todas las gorutinas)
	wg.Add(1000000)       //aca le decimos cuantas gorutinas lanzaremos. En este caso 1000000 por que por cada iteraccion del siguiente For nosotros lanzaremos una gorutina para saber si el numero es primo

	for i := 1; i <= 1000000; i++ {
		go func(a int) { // GOURUTINAS (APLICANDO CONCURRENCIA). La firma de la funcion anonima recibe un parametro a traves de la variable "a", esta es dada por la variable "i" (ver final de la funcion anonima al final de }: (i))
			defer wg.Done() //una vez terminado al gorutina lo quitaremos de la lista de espera "WaitGroup" con .Done()
			fmt.Println(a, esPrimo(a))
		}(i) // la funcion anonima recibe del exterior el valor de i, cuyo valor lo da el for

	}

	wg.Wait() //no continua con las siguientes instrucciones hasta que acaben todas los gorutines es decir hasta que la cola de gorutinas sync.WaitGroup este vacia
	fmt.Println("Termino el proceso")

}

func esPrimo(a int) (r bool) {
	r = true
	for i := 2; i <= a/2; i++ {
		if a%i == 0 {
			r = false
			return
		}
	}
	return
}

/*if esPrimo(i) {
	fmt.Println(i)
}*/
