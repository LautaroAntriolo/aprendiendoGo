// Se ingresan un conjunto de n alturas de personas por teclado. Mostrar la altura promedio de las personas.
package main

import "fmt"

func main(){
	var cantidad int
	var suma, valor float32
	
	fmt.Print("Ingrese el número de personas: ")
	fmt.Scan(&cantidad)
	for contador:= 0; contador < cantidad; contador++ {
		fmt.Print("Ingresá la altura de la persona: ")
		n, err:= fmt.Scan(&valor)
		if err != nil{
			fmt.Println("Error: ",err)
			return
		}
		fmt.Printf("Se ingresó %d valores \n",n)
		suma = suma + valor
	}
	fmt.Printf("El promedio es de: %f", suma/float32(cantidad))
}