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

4. Chat en Tiempo Real

    Descripción: Desarrolla una aplicación de chat simple utilizando WebSockets. Este proyecto te permitirá aprender sobre la comunicación en tiempo real y la gestión de conexiones concurrentes en Go.
    Recursos: Puedes utilizar el paquete github.com/gorilla/websocket para manejar WebSockets. Hay tutoriales disponibles que explican cómo construir un chat básico, como este tutorial sobre WebSockets en Go.

5. Scraper de Páginas Web

    Descripción: Crea un scraper que extraiga información de una página web específica (por ejemplo, títulos de artículos, precios de productos, etc.). Este proyecto te ayudará a familiarizarte con el manejo de solicitudes HTTP y el análisis de HTML.
    Recursos: Puedes usar el paquete golang.org/x/net/html para analizar documentos HTML. Un buen punto de partida es este video sobre cómo hacer scraping con Go: Scraping con Go.

6. Sistema de Autenticación

    Descripción: Implementa un sistema básico de autenticación y registro de usuarios. Este proyecto te permitirá trabajar con bases de datos, manejar contraseñas (hashing) y gestionar sesiones.
    Recursos: Puedes usar gorm para la interacción con bases de datos y bcrypt para el hashing de contraseñas. Un tutorial útil es este sobre autenticación en Go.
