package main

import "github.com/luistyle03/edcurso-go/07MetodosInterfacesDeferPanicRecover/0702Interfaces/animales"

func main() {
	var p animales.Perro
	var g animales.Gato
	p.Nombre = "Firulais"
	g.Nombre = "Bigotina"
	/*REALIZAR DE ESTA FORMA NO ES OPTIMA, MUCHO SE REPITE CODIGO POR ELLO USAMOS INTERFACES, QUE NOS PEMITE
	USAR UN SOLO METODO EN EL CUAL SU PARAMETRO MASCOTA ACEPTA PERRO O GATO
			AdoptarPerro(p)
			AdoptarGato(g)
	*/

	AdoptarMascota(p)
	AdoptarMascota(g)

}

/*
REALIZAR DE ESTA FORMA NO ES OPTIMA, MUCHO SE REPITE CODIGO POR ELLO USAMOS INTERFACES, QUE NOS PEMITE
USAR UN SOLO METODO EN EL CUAL SU PARAMETRO MASCOTA ACEPTA PERRO O GATO
func AdoptarPerro(p animales.Perro) {
	p.Comunicarse()
}

func AdoptarGato(g animales.Gato) {
	g.Comunicarse()
}
*/
func AdoptarMascota(m animales.Mascota) {
	m.Comunicarse()
}
