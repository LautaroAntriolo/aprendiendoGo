/*
4.Crear un programa que simule el comportamiento de un reloj digital, escribiendo en
formato de horas, minutos y segundos, iniciando desde las 00:00:00 horas hasta las
23:59:59 horas
*/
package ejemplosm4_8

import "fmt"

func RelojDigital(){
	horario := make([]int, 3)
	fmt.Print(horario)
	for i:=0; i<24; i++{
		horario[0] = i
		for i:=0; i<60; i++{
			horario[1] = i
			for i:=0; i<60; i++{
				horario[2] = i
				fmt.Print(horario)
			}
			fmt.Print(horario)
		}
		fmt.Print(horario)
	}
}