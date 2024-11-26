package main
import (
	"fmt"
)

var numero11, numero22 int = 50,40

func suma_2(param1 int, param2 int) (x1,y2 int){
	x1 = param1 + param2 
	y2 = param1 - param2
	return 
}

func tiposDeDatos(param3 float64, param4 int) (float64) {
	x3 := param3 * float64(param4)
	return x3 
}

func main(){
	uno, dos := suma_2(numero11,numero22)
	fmt.Println(uno)
	fmt.Println(dos)
}
