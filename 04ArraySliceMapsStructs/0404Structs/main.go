package main

import "fmt"

//Serian como las clase en otros lenguajes de programacion
//Como esta en mayuscula el nombre de struct(Persona), entonces podra ser usado por otro paquetes
type Persona struct {
	Nombre string
	Edad   uint8
}

type MejorPersona struct {
	Nombre string
	Edad   uint8
	correo []string
}

func main() {

	//Declarando y luego asignando struct
	var persona1 Persona
	persona1.Nombre = "Pedro"
	persona1.Edad = 34

	fmt.Println(persona1)
	fmt.Println(persona1.Nombre)

	//Otra forma de declarar y luego asignando struct
	persona2 := Persona{}
	persona2.Nombre = "Rosa"
	persona2.Edad = 71
	fmt.Println(persona2)

	//Declarando y asignando asignando struct (Forma normal)
	persona3 := Persona{Nombre: "Juan",
		Edad: 72,
	}
	fmt.Println(persona3)

	//Declarando y asignando asignando struct (Forma reducida)
	persona4 := MejorPersona{"Gaspar",
		20,
		[]string{
			"a@b.com",
			"c@d.com",
		}}
	fmt.Println(persona4)

	//Creando un slice de structs(forma larga)
	var personasss []Persona
	personasss = append(personasss, Persona{"Pedro", 30})
	personasss = append(personasss, Persona{"Juan", 50})
	fmt.Println(personasss)

	//Creando un slice de struct (forma corta)
	personasss2 := []Persona{
		Persona{"Rosa", 30},
		Persona{"gasparucho", 50},
	}

	fmt.Println(personasss2)

}
