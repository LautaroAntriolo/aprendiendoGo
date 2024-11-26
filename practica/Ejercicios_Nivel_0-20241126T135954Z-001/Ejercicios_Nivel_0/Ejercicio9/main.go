//En una empresa trabajan n empleados cuyos sueldos oscilan entre $100 y $500,
// realizar un programa que lea los sueldos que cobra cada empleado e informe cuántos 
// empleados cobran entre $100 y $300 y cuántos cobran más de $300. Además el programa
// deberá informar el importe que gasta la empresa en sueldos al personal.
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func empleados() (float32, int, int) {
    
	var entrada string
	var sueldoTotales float32
    var masQueElPromedio int
    var menosQueElPromedio int
    
    for {
        fmt.Print("Ingresá tu sueldo o 'n' para terminar: ")
        fmt.Scan(&entrada)
        if strings.ToLower(entrada) == "n" {
            break
        }
        sueldo, err := strconv.ParseFloat(entrada, 32)
        if err != nil {
            fmt.Println("Error al leer el sueldo")
            return 0, 0, 0
        }
        sueldoTotales += float32(sueldo)
        if sueldo >= 100 && sueldo <= 300 {
            menosQueElPromedio++
        } else if sueldo > 300 {
            masQueElPromedio++
    }}
    return sueldoTotales, menosQueElPromedio, masQueElPromedio}

func main() {
    sueldoTotales, menosQueElPromedio, masQueElPromedio := empleados()
    fmt.Println("Sueldo total:", sueldoTotales)
    fmt.Println("Menos que el promedio:", menosQueElPromedio)
    fmt.Println("Mas que el promedio:", masQueElPromedio)
}