package main

import (
	//"practica/fechas"
	//"practica/lecturas"
	"practica/programacionEstructuradaEnGo/ejemplosM1"
	"practica/programacionEstructuradaEnGo/ejemplosM2"
	"fmt"
)

func main(){
	//fechas.Edad(1996,12,3)
	//fechas.TiempoANavidad();
	// Recordatorio(anio int, mes time.Month, dia int, hora int, minutos int, mensaje string){
	//fechas.Recordatorio(2024,11,17,20,58,"A ver si esto funciona!")
	//datos.Lecturas("/home/lautaro/Documentos/learningGo/practica/empleados_actualizados.csv")

	// ejemplosM1.Volumenes("cono",4)
	// ejemplosM1.NumeroMaGrande(4)
	// ejemplosM1.Calculadora()
	// ejemplosM1.Teatro(15)
	// ejemplosM1.Descuento(1500,"Carnes")
	// ejemplosM1.CalcularTotales()
	ejemplosM1.DescomponerNumero(987)
	ejemplosM1.DescomponerEInvertir(987)
	fmt.Println("")
	ejemplosM2.MayoNumero(1,2,3,4,5)
	fmt.Print(ejemplosM2.AnioBisiesto(1992)) // true 
	fmt.Print(ejemplosM2.AnioBisiesto(2000)) // true
	fmt.Print(ejemplosM2.AnioBisiesto(1900)) // false
	
}
