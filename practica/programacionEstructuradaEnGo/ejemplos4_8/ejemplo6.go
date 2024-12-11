package ejemplosm4_8

import "fmt"

/*
6.Generar los primeros diez números perfectos.

Nota: Un Número perfecto es aquel que es igual a la suma de sus divisores,
excluyéndose el propio número. Por ejemplo, el número 28 presenta 5 divisores
menores y distintos de 28, que son: 1, 2, 4, 7 y 14. Al sumarlos da como resultado 28*/

func esPerfecto(numero int) bool {
	suma := 0
	for i := 1; i < numero; i++ {
		if numero%i == 0 {
			suma += i
		}
	}
	return suma == numero
}

func NumerosPerfectos(cantidad int) []int {
	numerosPerfectos := []int{}
	numero := 1

	for len(numerosPerfectos) < cantidad {
		if esPerfecto(numero) {
			numerosPerfectos = append(numerosPerfectos, numero)
		}
		numero++
	}
	fmt.Print(numerosPerfectos)

	return numerosPerfectos
}