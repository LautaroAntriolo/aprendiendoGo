package ejemplosm3

import (
	"math"
	"fmt"
)

// 3.Crear un programa que calcule si un número es primo o no.

func NumeroPrimo(numero int) bool{
	var esPrimo bool=true
	if numero <= 1 {
		fmt.Printf("\nEl número %v NO es primo", numero)
		return false
	}
	numeroFinal := int(math.Sqrt(float64(numero)))
	// i no es 1, porque todos los números son divisibles por 1
	for i := 2; i <numeroFinal+1; i++ {
		if numero%i == 0{
			esPrimo = false
			fmt.Printf("\nEl número %v NO es primo", numero)
			return esPrimo
		}
	}
	fmt.Printf("\nEl número %v es primo", numero)
	return esPrimo
}