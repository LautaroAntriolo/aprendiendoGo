/*
7.Según historias cuentan que el inventor del juego de ajedrez, el rey que le solicitó
hacerlo le preguntó cómo quería que este le pagase, a lo cual el inventor le contesto
que ubicara un grano de trigo en el primer cuadro del tablero, en el segundo ubicara
el doble del primero, en el tercero el doble del segundo y así sucesivamente hasta
completar los 64 cuadrados.
Se debe crear un programa que muestre un tablero (Puede ser sin líneas) en donde se
indique el número de granos de trigo debió colocar el rey por cada cuadro.

Nota: El ejemplo es ilustrativo, se debe basar en el tablero de ajedrez.

1 2 3 4 5 6
12 11 10 9 8 7
13 14 15 16 17 18
24 23 22 21 20 19
25 26 27 28 29 30
36 35 34 33 32 31
37 38 39 40 41 42
*/
package ejemplosm4_8

import "fmt"
func GranosDeTrigo(){
	var contador int = 0
	var estadoUno int = 0
	var estadoDos int = 0
	for  contador<6{
		estadoUno++
		fmt.Print(estadoUno)
		estadoDos++
		estadoUno+=estadoDos
		contador++
	}
}