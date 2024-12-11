package ejemplosm3

import (
	"fmt"
)

// 1.Crear un programa que reciba una serie de números y calcule la sumatorio de todos estos.
func Sumatoria(numeros ...int) int{
	resultado := 0
	for _, num := range numeros{
		resultado+=num
	}
	fmt.Printf("El resultado de la suma es %v",resultado)
	return resultado
}