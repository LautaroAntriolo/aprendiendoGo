#En una empresa trabajan n empleados cuyos sueldos oscilan entre $100 y $500,
# realizar un programa que lea los sueldos que cobra cada empleado e informe cuántos 
# empleados cobran entre $100 y $300 y cuántos cobran más de $300. Además el programa
# deberá informar el importe que gasta la empresa en sueldos al personal.

def empleados():
    sueldoTotales = 0
    entrada = input("Ingresá tu sueldo o n para terminar")
    masQueElPromedio = 0
    menosQueElPromedio = 0
    while entrada != "n":
        entrada = input("Ingresá tu sueldo o n para terminar")
        sueldoTotales+=entrada
        if entrada >= 100 and entrada <=300:
            menosQueElPromedio+=1
        elif entrada > 300:
            masQueElPromedio+=1
    return sueldoTotales, menosQueElPromedio, masQueElPromedio
            