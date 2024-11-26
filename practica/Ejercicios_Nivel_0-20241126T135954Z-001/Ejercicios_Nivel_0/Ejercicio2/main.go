// Mostrar todos los múltiplos de 8 que hay hasta el valor 500.
// Debe aparecer en pantalla 8 - 16 - 24, etc
package main
import "fmt"

func main(){

	var numero int = 0
	var lista []int
	for numero < 500{
		if numero%8 == 0{
			lista = append(lista, numero)
			numero++
		}
		numero++
	}
	fmt.Println(lista)
}