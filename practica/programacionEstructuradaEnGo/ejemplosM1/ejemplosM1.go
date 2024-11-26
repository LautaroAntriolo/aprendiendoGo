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
func Calculadora() {
    var resultado, numero float64
    var operacion string
    continuar := "s"
    
    fmt.Print("Ingresá el primer número: ")
    fmt.Scan(&resultado)
    
    for continuar == "s" {
        fmt.Print("Ingresá la operación (+, -, *, /): ")
        fmt.Scan(&operacion)
        
        fmt.Print("Ingresá el siguiente número: ")
        fmt.Scan(&numero)
        
        // Realizar la operación
        switch operacion {
        case "+":
            resultado += numero
        case "-":
            resultado -= numero
        case "*":
            resultado *= numero
        case "/":
            if numero != 0 {
                resultado /= numero
            } else {
                fmt.Println("Error: No se puede dividir por cero")
                continue
            }
        default:
            fmt.Println("Operación no válida")
            continue
        }
        
        fmt.Printf("Resultado actual: %.2f\n", resultado)
        
        fmt.Print("Querés continuar? s/n: ")
        fmt.Scan(&continuar)
        continuar = strings.ToLower(continuar)
    }
    
    fmt.Printf("Resultado final: %.2f\n", resultado)
}
/*4. En un teatro cada cliente paga $10000 por entrada y cada función le cuesta al teatro $300000 por atención prestada.
Por cada cliente que entre al teatro debe pagar un costo de 2000 por aseo. Desarrolar un programa que reciba
el numero de clientes de una función y devuelva el valor de las ganancias obtenidas*/

func Teatro(clientes int){
    
    const valorTeatro = 300000
    
    var aseo int = 2000

    var gananciaParcial float32 = float32(clientes) * float32(10000 - aseo)

    var gananciaTotal float32 = valorTeatro - gananciaParcial

    fmt.Printf("Las ganancias totales son: %.2f\n", gananciaTotal)

}

/*5.En un supermercada se ofrecen descuentos por el total del valor en cada uno de los sigueintes productos: 
carnes: 10%
frutas: 5%
aseo: 7%
dulces: 9%
*/

func Descuento(precio float32, producto string) float32{
    var resultado float32
    switch strings.ToUpper(producto) {
    case "CARNES":
        resultado = precio*0.9
    case "FRUTAS":
        resultado = precio*0.95
    case "ASEO":
        resultado = precio*0.93
    case "DULCES":
        resultado = precio*0.91
    }
    fmt.Printf("El resultado es: %.2f\n", resultado)
    return resultado
}

/*6.Una persona desea llevar $10.550 en dulces, $50.000 en carne y $35.000 en
productos de aseo. Se necesita calcular el descuento total por cada tipo de producto,
posteriormente entregar el valor total a pagar con descuento y sin descuento.*/


func CalcularTotales() {
	productos := make(map[string]float32)

	fmt.Println("Ingrese los productos y sus precios. Escriba 'FIN' para terminar.")
	for {
		var producto string
		var precio float32

		fmt.Print("Producto: ")
		fmt.Scanln(&producto)

		if strings.ToUpper(producto) == "FIN" {
			break
		}

		fmt.Print("Precio: ")
		_, err := fmt.Scanln(&precio)
		if err != nil || precio < 0 {
			fmt.Println("Precio inválido. Intente de nuevo.")
			continue
		}

		productos[producto] = precio
	}

	var descuentoTotal, totalSinDescuento, totalConDescuento float32

	fmt.Println("\nCálculo de descuentos y totales:")
	for producto, precio := range productos {

		descuento := Descuento(precio, producto)
		precioConDescuento := precio - descuento

		descuentoTotal += descuento
		totalSinDescuento += precio
		totalConDescuento += precioConDescuento

		// detalle de cada producto
		fmt.Printf("%s:\n", strings.ToUpper(producto))
		fmt.Printf("  Precio original: $%.2f\n", precio)
		fmt.Printf("  Descuento aplicado: $%.2f\n", descuento)
		fmt.Printf("  Precio con descuento: $%.2f\n\n", precioConDescuento)
	}

	fmt.Println("Resultado:")
	fmt.Printf("  Total sin descuento: $%.2f\n", totalSinDescuento)
	fmt.Printf("  Total descuentos: $%.2f\n", descuentoTotal)
	fmt.Printf("  Total con descuento: $%.2f\n", totalConDescuento)
}
/*7 Verificar si dos números ingresados son iguales*/
func IgualdadDeDosNumeros(numero1 float32, numero2 float32) bool {
	primero := float64(numero1) 
	segundo := float64(numero2) 
	var estado bool = false
	if primero == segundo{
		estado = true
	}

	return estado
}
/*8 VErificar si un número es menor a 10*/
func NumeroMenorA(numero float32, tope int) bool {
	var estado bool = false
	if numero < float32(tope){
		estado = true
	}else if numero == float32(tope){
		fmt.Print("los números son iguales. Volvé a ingresarlos")
	}
	
	return estado
}
/*9 Hacer un programa que determine verdadero o falso si un número es mayor o igual que 10 y menor que 20
o mayor que 30*/
func RangoDeNumeros(numero int) bool {
	var estado bool = false
	if numero >= 10 && numero < 20 {
		estado = true
	}else if numero > 30{
		estado = true
	}
	return estado
}
/*10 Dado un número de 3 cifras, descomponerlo en sus unidades fundamentales y decir la cantidad
correspondiente. Ejemplo: el número 631 tiene una unidad, 3 decenas y 6 centesimas*/
func DescomponerNumero(numero int16){
	centenas := numero/100
	decenas := (numero / 10) % 10
	unidades := numero % 10

	fmt.Printf("Número: %d\n", numero)
	fmt.Printf("Centenas: %d\n", centenas)
	fmt.Printf("Decenas: %d\n", decenas)
	fmt.Printf("Unidades: %d\n", unidades)
}
/*11 Del ejercicio anterior, descomponer un número dado e invertirlo. Por ejemplo: 582 pasaría a ser 285*/

func DescomponerEInvertir(numero int16) int16 {
	centenas := numero/100
	decenas := (numero / 10) % 10
	unidades := numero % 10
	
	invertido := (unidades * 100) + (decenas * 10) + centenas

	fmt.Printf("Número original: %d, Número invertido: %d\n", numero, invertido)
	return invertido
}