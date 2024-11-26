package main

import "fmt"

func operaciones()(resultado float32){
    x :=1
    var valor float32

    for x<=5{
        fmt.Print("Ingresá un valor:")
        fmt.Scan(&valor)
        resultado = resultado + valor
        x++
    }
    fmt.Print(resultado)
    return resultado
}

func main() {
    var horasTrabajadas int
    var costoHora float32
    var sueldo float32

    fmt.Print("Ingrese las horas trabajadas por el empleado:")
    fmt.Scan(&horasTrabajadas)
    fmt.Print("Ingrese el pago por hora:")
    fmt.Scan(&costoHora)
    sueldo=float32(horasTrabajadas) * costoHora
    fmt.Print("El sueldo total del operario es ",sueldo)
    operaciones()
}
