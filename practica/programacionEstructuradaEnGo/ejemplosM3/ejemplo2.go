package ejemplosm3

import (
	"fmt"
)
/* 2.Escribir una función que calcule cuantos números pares hay comprendidos entre 
dos números límites (sin incluirlos).*/
func ParesEntreDosNumeros(inicio int, fin int) (int,int ){
	var paresCant int = 0
	var imparesCant int = 0
	
	for i := inicio; i < fin-1; i++ {
		if i%2==0{
			paresCant++
		}else{
			imparesCant++
		}
	}
	fmt.Printf("\nEntre %v y %v hay %v pares y %v impares", inicio,fin,paresCant,imparesCant)
	return paresCant, imparesCant
}