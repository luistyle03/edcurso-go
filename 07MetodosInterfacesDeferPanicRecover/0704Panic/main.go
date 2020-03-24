package main

import "fmt"

func main() {
	division(3, 0)
}

func division(dividendo int64, divisor int64) {
	fmt.Println("1")
	defer fmt.Println("Esto se ejecutara pase lo que pase")
	fmt.Println("2")
	if divisor == 0 {
		panic("No se puede dividir entre cero")
	}
	fmt.Println(dividendo / divisor)
}
