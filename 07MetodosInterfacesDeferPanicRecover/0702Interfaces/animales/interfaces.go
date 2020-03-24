package animales

type Mascota interface {
	/*
		//Como se implementa la interfaz? Simplemente los metodos de la interfaz tienen que tener el
		mismo nombre que el metodo de cada una de las clases(las instrucciones del metodo de la clase
		es su implementacion)
		EL METODO DE LA CLASE NO PUEDE RECIBIR UN PUNTERO:
		func (g *Gato) Comunicarse() ESTO NO ES POSIBLE
		PERO SI PUEDE TENER COMO PARAMETROS PUNTEROS:
		func (g Gato) Comunicarse(parametro1 *int32) ESTO NO ES POSIBLE
	*/
	Comunicarse()
}
