package ejemplos5_1 

import "fmt"

type NotaPorMateria struct {
	materia string
	nota    float32
}

type Estudiante struct{
	nombre string
	edad   int
	notas  []*NotaPorMateria
}
func ingresarEstudiante() int{
	var Estu = Estudiante{}
	var nombre string
	var edad int
	fmt.Print("Ingrese el nombre del alumno: ")
	fmt.Scan(&nombre)
	fmt.Print("Ingrese la edad del alumno: ")
	fmt.Scan(&edad)
	Estu.nombre = nombre
	Estu.edad = edad
	return edad
}

func FinDeAnio(){
	ingresarEstudiante()
	// fmt.Printf("El nombre del alumno es %s y tiene %v años", )
}