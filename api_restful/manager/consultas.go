package manager


import (
    "net/http"
    "strconv"
    "github.com/labstack/echo/v4"
)

type Task struct {
	ID          int `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool `json:"done"`
}

var tasks []Task

// tomar todas las tareas
func GetTasks(c echo.Context) error {
    return c.JSON(http.StatusOK, tasks)
}

// Crear una nueva tarea
func CreateTask(c echo.Context) error {
	task := new(Task)
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Error al procesar los datos",
		})
	}

	// Asignamos ID
	if len(tasks) > 0 {
		task.ID = tasks[len(tasks)-1].ID + 1
	} else {
		task.ID = 1
	}

	task.Done = false // Las tareas comienzan sin completar
	tasks = append(tasks, *task)

	return c.JSON(http.StatusCreated, task)
}


// Actualizar tarea
func UpdateTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID inválido",
		})
	}

	for i, t := range tasks {
		if t.ID == id {
			updatedTask := new(Task)
			if err := c.Bind(updatedTask); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "Error al procesar los datos",
				})
			}

			// Actualizar solo los campos relevantes
			tasks[i].Title = updatedTask.Title
			tasks[i].Description = updatedTask.Description
			tasks[i].Done = updatedTask.Done
			return c.JSON(http.StatusOK, tasks[i])
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "Tarea no encontrada",
	})
}


// Eliminar tarea
func DeleteTask(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            return c.JSON(http.StatusOK, map[string]string{"message": "Task deleted"})
        }
    }
    return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
}