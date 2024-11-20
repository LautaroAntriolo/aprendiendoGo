package ejemplosM1

/* Crear un programa que permita calcular el volumen de un cilindro, una esfera y 
un cono. Las ecuaciones de las respecticas figuras son: 

cilindro = pi * r^2
esfera = 4/3 * pi * r^4
cono = 1/3 * pi * r^2

*/
import(
	"fmt"
	"strings"
	"math"
)
func redondear(valor float64, decimales int) float64 {
	/*Creo esta función para redondear el número, y no la salida.*/
	factor := math.Pow(10, float64(decimales))
	return math.Round(valor*factor) / factor
}

func Volumenes(figura string, radio float64){
	figura = strings.ToLower(figura) 
	if figura == "cilindro"{
		//resultado = math.Pi * radio* radio // no existe exponencial en Go. :(
		resultado := math.Pi * math.Pow(radio,2)
		fmt.Printf("El volumen del cilindro es de %.2f\n", redondear(resultado,2))
	}else if figura == "esfera"{
		resultado := 4/3 * math.Pi * math.Pow(radio,4)
		fmt.Printf("El volumen de la esfera es de %.2f\n", redondear(resultado,2))
	}else if figura == "cono"{
		/* resultado := 1/3  da 0, porque es división entera. 
		Se usa float64 para divisiones que deben conservar decimales.*/
		var a, b float64 = 1, 3
		resultado := a/b * math.Pi * math.Pow(radio,2)
		fmt.Printf("El volumen del cono es de %.2f\n", redondear(resultado,2))
	}
}
