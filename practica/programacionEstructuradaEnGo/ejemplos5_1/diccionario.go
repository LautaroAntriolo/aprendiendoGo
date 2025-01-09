package ejemplos5_1

import (
	"Strings"
	"fmt"
)

func cargar(diccionario map[string]string) {

	var espaniol string
	var ingles string
	var opcion string

	for {
		fmt.Print("Ingrese una palabra en ingles: ")
		fmt.Scan(&ingles)
		fmt.Print("Ingrese su traducción en español: ")
		fmt.Scan(&espaniol)
		diccionario[ingles] = espaniol
		fmt.Print("Desea ingresar otra palabra? ")
		fmt.Scan(&opcion)
		if strings.ToLower(opcion) == "n" || strings.ToLower(opcion) == "no" {
			break
		}
	}
}
func mostrarDiccionario(diccionario map[string]string) {
	fmt.Printf("\n")
	for clave, valor := range diccionario {
		fmt.Printf("%s: %s \n", clave, valor)
	}
}
func borrarDelDiccionario(diccionario map[string]string) {
	var palabraEnIngles string
	fmt.Print("Ingrese una palabra en ingles que quiera borrar del diccionario: ")
	fmt.Scan(&palabraEnIngles)
	_, existe := diccionario[palabraEnIngles]
	if existe {
		delete(diccionario, palabraEnIngles)
		fmt.Printf("Se ah eliminado '%s' del diccionario", palabraEnIngles)
	} else {
		fmt.Printf("No existe %s ", palabraEnIngles)
	}
}
func actualizarPalabra(diccionario map[string]string) {
	var palabraEningles string

	fmt.Print("Que palabra quiere modificar? ")
	fmt.Scan(&palabraEningles)
	_, existe := diccionario[palabraEningles]

	if existe {

		var opciones int
		var PreTraduccionEspanol string
		var traduccionEspanol string
		var aux string

		fmt.Print("Que quiere modificar? \n 1.Palabra en ingles \n 2. palabra en español \n modificar por una nueva palabra \n 4. salir ")
		fmt.Scan(&opciones)
		switch opciones {
		case 1:
			PreTraduccionEspanol = diccionario[palabraEningles]
			aux = palabraEningles
			delete(diccionario, palabraEningles)
			fmt.Println("Ingrese la nueva palabra en Ingles")
			fmt.Scan(&palabraEningles)
			diccionario[palabraEningles] = PreTraduccionEspanol
			fmt.Printf("La palabra %v ha sido cambiada por %v \n", aux, palabraEningles)
			mostrarDiccionario(diccionario)
		case 2:
			fmt.Println("Ingrese la nueva traducción en Español")
			fmt.Scan(&traduccionEspanol)
			aux = diccionario[palabraEningles]
			diccionario[palabraEningles] = traduccionEspanol
			fmt.Printf("La palabra %v ha sido cambiada por %v\n", aux, palabraEningles)
			mostrarDiccionario(diccionario)
		case 3:
			delete(diccionario, palabraEningles)
			fmt.Println("Ingrese la nueva palabra en Ingles")
			fmt.Scan(&palabraEningles)
			fmt.Println("Ingrese la nueva traducción en Español")
			fmt.Scan(&traduccionEspanol)
			diccionario[palabraEningles] = traduccionEspanol
			fmt.Println("su nuevo diccionario es: ")
			mostrarDiccionario(diccionario)
		case 4:
			return
		}
	} else {
		fmt.Println("No existe")
	}

}
func Main() {
	var borrar string
	var actualizacion string
	diccionario := make(map[string]string)
	cargar(diccionario)
	mostrarDiccionario(diccionario)
	fmt.Printf("\n")
	fmt.Println("Quiere borrar alguna palabra del diccionario? ")
	fmt.Scan(&borrar)
	if strings.ToLower(borrar) == "s" || strings.ToLower(borrar) == "si" {
		borrarDelDiccionario(diccionario)
		mostrarDiccionario(diccionario)
	}
	fmt.Printf("\n")
	fmt.Println("Quiere realizar alguna actualización? ")
	fmt.Scan(&actualizacion)
	if strings.ToLower(actualizacion) == "s" || strings.ToLower(actualizacion) == "si" {
		actualizarPalabra(diccionario)
	}
}
