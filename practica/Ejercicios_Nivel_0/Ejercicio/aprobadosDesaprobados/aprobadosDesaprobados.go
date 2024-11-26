package aprobadosDesaprobados

import (
	"fmt"
)

// Escribir un programa que solicite ingresar 10 notas de alumnos
//  y nos informe cuántos tienen notas mayores o iguales a 7 y cuántos menores.
func PromedioDe10() {
    x := 10
    var contador int = 0
    var listaAprobados []float32
    var listaDesaprobados []float32

    for contador < x {
        var numero float32
        fmt.Print("Ingresá la nota de algún alumno: ")
        fmt.Scan(&numero)

        // Validamos que la nota esté entre 0 y 10
        if numero < 0 || numero > 10 {
            fmt.Println("Por favor ingrese una nota válida entre 0 y 10")
            continue // Volvemos a pedir el número sin incrementar el contador
        }else{
			if numero >= 4 {
				listaAprobados = append(listaAprobados, numero)
			} else {
				listaDesaprobados = append(listaDesaprobados, numero)
			}
			contador++
		}

    }

    totales := len(listaAprobados) + len(listaDesaprobados)
    
    fmt.Printf("\nResumen de notas:\n")
    fmt.Printf("Cantidad de aprobados: %d de %d (%.1f%%)\n", 
        len(listaAprobados), 
        totales, 
        float32(len(listaAprobados))/float32(totales)*100)
    fmt.Printf("Cantidad de desaprobados: %d de %d (%.1f%%)\n", 
        len(listaDesaprobados), 
        totales, 
        float32(len(listaDesaprobados))/float32(totales)*100)
    
    fmt.Printf("\nDetalle de notas:\n")
    fmt.Printf("Notas aprobadas: %v\n", listaAprobados)
    fmt.Printf("Notas desaprobadas: %v\n", listaDesaprobados)
}