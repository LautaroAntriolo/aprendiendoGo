// Realizar un programa que permita cargar dos listas de 15 valores cada una.
// Informar con un mensaje cual de las dos listas tiene un valor acumulado mayor
// (mensajes "Lista 1 mayor", "Lista 2 mayor", "Listas iguales")
// Tener en cuenta que puede haber dos o más estructuras repetitivas en un algoritmo.
package main

import (
	"fmt"
)

func listaDe15Valores(listaUno [15]int,listaDos [15]int){
	
	var sumListaUno int
	var sumListaDos int

	for _, value := range listaUno {
		sumListaUno+=value
    }
	for _, value := range listaDos {
		sumListaDos+=value
    }
	if sumListaUno>sumListaDos {
		fmt.Println("Lista 1 mayor")
	}else if sumListaUno<sumListaDos{
		fmt.Println("Lista 2 mayor")
	}else{
		fmt.Println("Listas iguales")
	}


}
func main(){
	array := [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,13, 14, 15}
	array2 := [15]int{11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,13, 14, 15}
    listaDe15Valores(array,array2)
}