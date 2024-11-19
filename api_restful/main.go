package main

import (
    "api_restful/fechas"
    "api_restful/manager"
    "github.com/labstack/echo/v4"
    "html/template"
    "io"
    "net/http"
)

// Template es un wrapper para el motor de plantillas de Echo
type Template struct {
    templates *template.Template
}

// Render es un método que permite a Echo renderizar plantillas
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
    e := echo.New()

    // Configurar el motor de plantillas
    t := &Template{
        templates: template.Must(template.ParseFiles("public/index.html")),
    }
    

    // Ruta principal que llama a la función TiempoANavidad
    e.GET("/", func(c echo.Context) error {
        // Llamar a la función TiempoANavidad que devuelve un mensaje
        tiempo := fechas.TiempoANavidad()
        
        // Pasar el valor a la plantilla HTML
        return c.Render(http.StatusOK, "index.html", map[string]interface{}{
            "TiempoANavidad": tiempo,
        })
    })
	e.Renderer = t

    // Rutas para tareas
    e.GET("/tasks", manager.GetTasks)
    e.POST("/tasks", manager.CreateTask)
    e.PUT("/tasks/:id", manager.UpdateTask)
    e.DELETE("/tasks/:id", manager.DeleteTask)

    // Servir archivos estáticos
    e.Static("/static", "public")

    // Iniciar servidor en el puerto 1323
    e.Logger.Fatal(e.Start(":1323"))
}
