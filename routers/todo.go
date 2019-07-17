package routers

import (
	"go-crud/controllers"
	"go-crud/services"

	"github.com/labstack/echo"
)

// TodoRouter ...
type TodoRouter struct {
	Echo     *echo.Echo
	Services TodoRouterServices
}

// TodoRouterServices ...
type TodoRouterServices struct {
	TodoService services.ITodoService
}

// New ...
func (r *TodoRouter) New(echo *echo.Echo) *echo.Echo {
	r.Echo = echo
	return r.init(r.Services)
}

func (r *TodoRouter) init(services TodoRouterServices) *echo.Echo {
	todoController := &controllers.TodoController{TodoService: services.TodoService}

	r.Echo.GET("/todos", todoController.AllTodos)
	r.Echo.GET("/todos/:id", todoController.ATodo)
	r.Echo.POST("/todos", todoController.CreateTodo)
	r.Echo.PUT("/todos/:id", todoController.UpdateTodo)

	r.Echo.GET("/test", todoController.TestHandler)

	return r.Echo
}
