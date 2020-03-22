package main

import "fmt"

func main() {
	n := []int8{3, 52, 4, 1, 89, 9, 100}
	maximo, minimo := maxymin(n)
	fmt.Printf("El maximo es %d y el minimo es %d \n", maximo, minimo)
}

func maxymin(numeros []int8) (max int8, min int8) {
	//tambien pudo haber funcionado asi:
	//func maxymin(numeros []int8) (int8, int8) {
	//var max, min int8
	for _, v := range numeros {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return
	// en caso haya sido esto: func maxymin(numeros []int8) (int8, int8) entonces el rerturn hubiera sido asi
	//return max, min
}
