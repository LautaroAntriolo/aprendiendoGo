2. API RESTful para Gestión de Tareas

Descripción:
Diseña una API RESTful básica que permita a los usuarios administrar una lista de tareas.

Requisitos:

    Usa el paquete net/http o frameworks como Gin o Echo.
    Implementa endpoints como:
        GET /tasks: Devuelve la lista de tareas.
        POST /tasks: Crea una nueva tarea.
        PUT /tasks/{id}: Marca una tarea como completada.
        DELETE /tasks/{id}: Elimina una tarea.
    Las tareas deben almacenarse en memoria (un slice de estructuras).

Objetivo:
Aprender a trabajar con servidores HTTP, JSON y CRUD básico.

3. Sistema de Gestión de Productos con Base de Datos

Descripción:
Construye un programa que administre productos en un inventario utilizando una base de datos SQLite.

Requisitos:

    Usa el paquete database/sql y un driver como github.com/mattn/go-sqlite3.
    Implementa las siguientes operaciones:
        Agregar un producto (nombre, precio, cantidad).
        Listar todos los productos.
        Actualizar la cantidad de un producto.
        Eliminar un producto.
    Opcional: añade menús en la terminal para interactuar con el usuario.

Objetivo:
Familiarizarse con bases de datos, SQL y operaciones básicas en Go.
