package controllers

import (
	"go-crud/services"
	"net/http"

	"github.com/labstack/echo"
)

// TodoController ...
type TodoController struct {
	TodoService services.ITodoService
}

// AllTodos ...
func (t *TodoController) AllTodos(ctx echo.Context) error {
	data, err := t.TodoService.GetAllTodos()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// ATodo ...
func (t *TodoController) ATodo(ctx echo.Context) error {
	todoID := ctx.Param("id")
	data, err := t.TodoService.GetATodo(todoID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, data)
}
