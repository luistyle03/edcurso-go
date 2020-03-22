package main

import "fmt"

func main() {

	//Declarando maps
	idiomas := make(map[string]string)
	//asignando maps
	idiomas["es"] = "Español"
	idiomas["en"] = "Ingles"
	idiomas["fr"] = "Frances"
	fmt.Println(idiomas)

	//Declarando y asignando mapas
	idiomas2 := map[string]string{
		"es": "Español",
		"en": "Ingles",
		"fr": "Frances",
	}
	fmt.Println(idiomas2)
	fmt.Println(idiomas2["en"])
	//Eliminando el valor con clave "en"
	delete(idiomas2, "en")
	fmt.Println(idiomas2)

	/*Verificando que un valor en el maps existe.Si el valor de maps con clave pt existe entonces idioma
	almacena el valor y ok es true, caso contrario ok es falso*/
	if idioma, ok := idiomas2["pt"]; ok {
		fmt.Println("Existe el maps con clave pt: ", idioma)
	} else {
		fmt.Println("NO Existe el maps con clave pt")

	}

}
