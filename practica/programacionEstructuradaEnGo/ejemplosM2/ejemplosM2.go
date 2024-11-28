package ejemplosM2

import (
	"fmt"
	"math"
	"sort"
	"strconv"
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
•5 decenas
•3 centenas
*/
func Descomponer(numero int){
	if numero <= 0 {
		fmt.Println("Ingresá un número positivo.")
		return
	}

	posiciones := []string{"unidades", "decenas", "centenas", "millares", "diezmillares", "cienmillares"}
	i := 0

	for numero > 0 {
		digito := numero % 10
		if digito > 0 {
			fmt.Printf("%d %s\n", digito, posiciones[i])
		}
		numero /= 10 
		i++;
	}
}

/*6.Construir una función que reciba un número entero y cuente el número de
veces que se repite un dígito (si existe).*/


func RepeticionNumero(numero int) map[int]int {
	
	if numero <0{
		fmt.Print("Ingresá un número válido. ")
	}
	
	var numeroString string = strconv.Itoa(numero)
	frecuencia := make(map[int]int)

	for _, num := range numeroString{
		// fmt.Println(num)

		digito := int(num - '0')
		// fmt.Println(digito)
		frecuencia[digito]++
	}
	fmt.Println(frecuencia)
	for digito, cantidad := range frecuencia {
		fmt.Printf("%d aparece %d veces\n", digito, cantidad)
	}
	return frecuencia
}

/*7.Crear una función que permita llenar un vector de n digitas (por teclado) y
calcule cual es el mayor de todos los elementos.*/

func EleccionNumeroMasGrande(numeros ...int) int{
	// Los múltiplas parámetros ya se representan como un slice.
	masGrande := numeros[0]
	
	if len(numeros) == 0 {
		return 0 
	}

	for _, numero := range numeros{
		if numero > masGrande{
			masGrande= numero
		}
	}
	fmt.Printf("el número mas grande es %d", masGrande)
	return masGrande
}
/*8.Crear una función que permita ordenar un vector de mayor a menor y
menor a mayor.*/
func OrdenarVectorOptimizado(vector []int, ascendente bool) []int {
	n := len(vector)
	for i := 0; i < n-1; i++ {
		swapped := false // genero esta variable que me permite cortar antes si es necesario.
		for j := 0; j < n-i-1; j++ {
			if ascendente && vector[j] > vector[j+1] {
				vector[j], vector[j+1] = vector[j+1], vector[j]
				swapped = true // Cambia a true si al recorrerlo hace algún cambio.
			}
			if !ascendente && vector[j] < vector[j+1] {
				vector[j], vector[j+1] = vector[j+1], vector[j]
				swapped = true
			}
		}
		// Si no hay cambios, entonces el slices ya está ordenada y salimos de la función.
		if !swapped {
			break
		}
	}
	fmt.Print(vector)
	return vector
}
func OrdenarVector(vector []int, ascendente bool) []int{

	if ascendente{
		sort.Ints(vector)
	}else{
		sort.Sort(sort.Reverse(sort.IntSlice(vector)))
	}
	// fmt.Print(vector)
	return vector
}

/*9.Construir un programa que, dados dos vectores, devuelva dos nuevos
vectores que junten los números pares e impares de cada uno.*/
func OrganizarVectoresEnParsEImpares(vector1, vector2 []int) ([]int, []int){

	var vectorPares []int
	var vectorImpares []int

	for _, numero := range vector1{
		if numero % 2 == 0{
			vectorPares = append(vectorPares, numero)
		}else{
			vectorImpares = append(vectorImpares, numero)
		}
	}
	for _, numero := range vector2{
		if numero % 2 == 0{
			vectorPares = append(vectorPares, numero)
		}else{
			vectorImpares = append(vectorImpares, numero)
		}
	}
	return vectorPares, vectorImpares
}

/*10. Crear un programa capaz de recibir una lista de datos y calcular la media,
mediana, moda y desviación estándar del conjunto de datos.*/
func mediaDeUnaLista(lista []int) float32 {
	var suma int 
	for _, num := range lista{
		suma += num
	}
	var resultado float64 = float64(suma) / float64(len(lista))
	return float32(resultado)
}
func medianaDeUnaLista(lista []int)float32{
	var posicionMedia int = len(lista)/2
	var listaOrdenada []int = OrdenarVector(lista, true)
	valorMediano := listaOrdenada[posicionMedia]
	return float32(valorMediano)
}
func modaDeUnaLista (lista []int)float32{

	organizacion := make(map[int]int)
	var moda int 
	var frecuanciaMaxima int

	for _, num := range lista{
		organizacion[num]++
	}

	for numero, frecuencia := range organizacion{
		if frecuencia > frecuanciaMaxima{
			moda = numero
			frecuanciaMaxima = frecuencia
		}
	}
	return float32(moda)

}
// Puede haber mas de una moda. EL código quedaría algo así: 
func modasDeUnaLista(lista []int) []int {
	
	organizacion := make(map[int]int)
	var frecuanciaMaxima int
	var modas []int

	for _, num := range lista {
		organizacion[num]++
	}

	for numero, frecuencia := range organizacion {
		if frecuencia > frecuanciaMaxima {
			frecuanciaMaxima = frecuencia
			modas = []int{numero}
		} else if frecuencia == frecuanciaMaxima {
			modas = append(modas, numero)
		}
	}

	return modas
}
func desviacionEstandar(lista []int) float32{

	tamanio := len(lista)
	if tamanio == 0 {
		return 0 
	}

	media := mediaDeUnaLista(lista)
	var sumaCuadrados float32

	for _, num := range lista {
		valor := float32(num) - media 
		sumaCuadrados += valor * valor
	}

	var desEst float64 = math.Sqrt(float64(sumaCuadrados) / float64(tamanio))

	return float32(desEst)
}

func Estadisticas(lista []int) (float32, float32, float32){
	
	var media float32 = mediaDeUnaLista(lista)
	fmt.Printf("La media es %.2f\n", media)
	
	var mediana float32 = medianaDeUnaLista(lista)
	fmt.Printf("la mediana es %.2f\n", mediana)

	var moda float32 = modaDeUnaLista(lista)
	fmt.Printf("la moda  es %.2f\n", moda)
	
	var desEstandar float32 = desviacionEstandar(lista)
	fmt.Printf("la desviacion estandar es %.2f\n",desEstandar)

	return media, mediana, moda

}

/*11.Construir una función que, dadas dos matrices, calcule la suma y
multiplicación de sus elementos. (Tener en cuenta propiedades de las matrices)*/


/*12.Construir una función en donde dada una matriz, calcule su matriz
transformada. (Tener en cuenta propiedades de las matrices)*/

/*13.Construir un programa que pueda recrear el juego Tic-tac-toe.Serán dos
jugadores posibles y gana quien alcance a completar 3 casillas en línea o diagonal,
por ejemplo:*/