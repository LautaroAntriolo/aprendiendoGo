import random
import time

# Leer los nombres del archivo de texto
with open("empleados.txt", "r", encoding="utf-8") as file:
    nombres = file.read().splitlines()  # Crear una lista de nombres, eliminando saltos de línea

# Archivo de salida donde se agregarán los datos
output_file = "empleados_actualizados.csv"

# Crear el archivo con la cabecera (si no existe)
with open(output_file, "w", encoding="utf-8") as output:
    output.write("Nombre,Apellido,Email,Teléfono\n")  # Escribir la cabecera

# Procesar cada nombre y agregar los datos al archivo de salida
for nombre in nombres:
    # Calcular los valores
    apellido = f"{nombre[2:]}{nombre[0] + nombre[2] + nombre[1]}" if len(nombre) > 2 else "N/A"
    mail = f"{nombre.lower()}{nombre[0].lower() + nombre[2].lower() if len(nombre) > 2 else ''}@gmail.com"
    pais = ["54","59","34"]
    fijo = ["2325","2334","11"]
    aleatorio = str(random.randint(1000, 9999))
    telefono = f"+{random.choice(pais)}{random.choice(fijo)}{aleatorio}"

    datos_persona = f"{nombre},{apellido},{mail},{telefono}\n"
    with open(output_file, "a", encoding="utf-8") as output:
        output.write(datos_persona)

    time.sleep(2)
    print(f"Datos agregados al archivo '{output_file}' correctamente.")
