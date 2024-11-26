// Desarrollar un programa que permita cargar n números enteros y luego 
// nos informe cuántos valores fueron pares y cuántos impares.
package main 

import "fmt"

func main(){
	var n, numero int 
	var cantidadPares,cantidadImpares int
	
	fmt.Print("ingresá un número entero: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		
		fmt.Print("Ingresá un número entero: ")
		fmt.Scan(&numero) 

		if numero%2 == 0 {
			cantidadPares++
		} else {
			cantidadImpares++
		}
	}

	fmt.Println("Cantidad de números pares:", cantidadPares)
	fmt.Println("Cantidad de números impares:", cantidadImpares)
}