package ejemplosM1

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

/* Crear un programa que permita calcular el volumen de un cilindro, una esfera y 
un cono. Las ecuaciones de las respecticas figuras son: 

cilindro = pi * r^2
esfera = 4/3 * pi * r^4
cono = 1/3 * pi * r^2

*/
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

/*2. Calcular el número mas grande de 5 números dados*/
func NumeroMaGrande(peticiones int){
	var numeroMasGrande uint64
	for i := 0; i<=peticiones; i++ {
		var numero int
		fmt.Println("Ingresá un número: ")
		fmt.Scan(&numero)
		if numero>int(numeroMasGrande){
			numeroMasGrande = uint64(numero)
		}
	}
	fmt.Println("El número mas grande es",numeroMasGrande)
}
/*3. REalziar las operaciones básicas con números de tipo entero y números de tpo flotante*/
/*4. En un teatro cada cliente paga $10000 por entrada y cada función le cuesta al teatro $300000 por atención prestada.
Por cada cliente que entre al teatro debe pagar un costo de 2000 por aseo. Desarrolar un programa que reciba
el numero de clientes de una función y devuelva el valor de las fanancias obtenidas*/
/*5.En un supermercada se ofrecen descuentos por el total del valor en cada uno de los sigueintes productos: 
carnes: 10%
frutas: 5%
aseo: 7%
dulces: 9%
*/
/*6. Una persona desea llevar 10550 en dulces, 50000 en carne y 35000 en productos de aseo. Se necesita calcular el descuento
titak oir cada tuoi de producto, posteriormente entregar el valor total a pagar con descuento y sin descuento*/

/*7 Verificar si dos números ingresados son iguales*/
/*8 VErificar si un número es menor a 10*/
/*9 Hacer un programa que determine verdadero o falso si un número es mayor o igual que 10 y menor que 20
o mayor que 30*/
/*10 Dado un número de 3 cifras, descomponerlo en sus unidades fundamentales y decir la cantidad
correspondiente. Ejemplo: el número 631 tiene una unidad, 3 decenas y 6 centesimas*/
/*11 Del ejercicio anterior, descomponer un número dado e invertirlo. Por ejemplo: 582 pasaría a ser 285*/