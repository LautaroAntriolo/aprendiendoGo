package ejemplosM2

import (
	"fmt"
)
/*1.Del capítulo “2. Estructuras de un programa en GO”, desarrollar los ejercicios
propuestos en modo de funciones. HECHO*/
/*2.Crear un programa que determine el mayor de 3, 4 y 5 números.*/
func MayoNumero(numeros ...int) float32{
 	
	var resultado float32

	for _, numero :=range numeros{
		if float32(numero) > resultado{
			resultado = float32(numero)
		}

	}
	fmt.Printf("el número mas grande es %.2f\n", resultado)
	return resultado
}

/*3.Se necesita crear un programa que reciba la fecha de un año e indique si
este es o no bisiesto*/
func AnioBisiesto(anio int) bool{
	if anio%4 == 0 && (anio%100 != 0 || anio%400 == 0) {
        return true
    }
    return false
}
/*4.Crear un programa que, dado un intervalo de números entero, verifique si la
suma de los elementos comprendidos en el intervalo da como resultado un
numero primo.*/

func ResultadoPrimo(fin int) bool {
	
	var datos = make([]int, fin)
	var resultado int

	for i := 0; i < fin; i++ {
		datos[i] = i + 1
	}

	for _, numero := range datos {
		resultado += numero
	}

	if resultado <= 1 {
		return false
	}
	for i := 2; i*i <= resultado; i++ {
		if resultado%i == 0 {
			return false
		}
	}
	return true
}
/*5.Construir un programa que reciba un numero de n dígitos y devuelva
(según el número de dígitos) las unidades utilizadas con su cantidad. Por ejemplo:
Se recibe 356, se debe retornar como respuesta:
6 unidades

42

•5 decenas
•3 centenas*/

/*6.Construir una función que reciba un número entero y cuente el número de
veces que se repite un dígito (si existe).*/
/*7.Crear una función que permita llenar un vector de n digitas (por teclado) y
calcule cual es el mayor de todos los elementos.*/
/*8.Crear una función que permita ordenar un vector de mayor a menor y
menor a mayor.*/
/*9.Construir un programa que, dados dos vectores, devuelva dos nuevos
vectores que junten los números pares e impares de cada uno.*/
/*10. Crear un programa capaz de recibir una lista de datos y calcular la media,
mediana, moda y desviación estándar del conjunto de datos.*/
/*11.Construir una función que, dadas dos matrices, calcule la suma y
multiplicación de sus elementos. (Tener en cuenta propiedades de las matrices)*/
/*12.Construir una función en donde dada una matriz, calcule su matriz
transformada. (Tener en cuenta propiedades de las matrices)*/

/*13.Construir un programa que pueda recrear el juego Tic-tac-toe.Serán dos
jugadores posibles y gana quien alcance a completar 3 casillas en línea o diagonal,
por ejemplo:*/