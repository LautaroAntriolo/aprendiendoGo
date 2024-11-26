package main

import (
	"fmt"
	"math"
)
// func suma(x int, y int) float32{
// 	resultado :=  x + y 
// 	return float32(resultado)
// }

func suma(x ,y int) float32{ // Como ambos parámetros serán del mismo tipo de dato, entonces le agregamos el tipo de dato al final.
	resultado :=  x + y 
	return float32(resultado)
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum float32) (x, y float32) {
	x = sum * float32(1.1)
	y = x - sum
	return // Esto se conoce como retorno "desnudo" deben usarse solo en funciones cortas
}

var numero1, numero2  int = 1, 2

func main(){
	fmt.Println("Grande bocaa! ")
	fmt.Println(math.Sqrt(5))
	fmt.Println(suma(numero1,numero2))

	a, b := swap("hola", "mundo")
	fmt.Println(a, b)
	
	fmt.Println(split(17))
}
