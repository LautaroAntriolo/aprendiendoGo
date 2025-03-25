2. API RESTful para Gestión de Tareas

Descripción:
Una API RESTful básica que permita a los usuarios administrar una lista de tareas.

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
