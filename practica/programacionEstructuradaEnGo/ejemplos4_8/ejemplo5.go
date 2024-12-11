package ejemplosm4_8

import "fmt"

func TablaDel(numero int) {
	fmt.Printf("Tabla del %v", numero)
	for i := 0; i<11; i++{
		fmt.Printf("\n%v * %v = %v",numero,i,numero*i)
	}
}