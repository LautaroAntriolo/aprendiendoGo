// Realizar un programa que imprima 25 términos de la serie 11 - 22 - 33 - 44, etc.
// (No se ingresan valores por teclado);
package main

import (
	"fmt"
)

func main(){
	fmt.Println("INICIO") 
	for i := 0; i < 25; i++ {
		fmt.Printf("%d%d - ",i,i)
	}
	fmt.Println("FIN") 
}