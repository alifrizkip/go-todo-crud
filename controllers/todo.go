package controllers

import (
	"fmt"
	"go-crud/requests"
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

// CreateTodo ...
func (t *TodoController) CreateTodo(ctx echo.Context) error {
	newTodoReq := new(requests.NewTodo)

	err := ctx.Bind(newTodoReq)
	if err != nil {
		return err
	}

	err = t.TodoService.CreateTodo(newTodoReq)
	if err != nil {
		return err
	}

	resp := struct {
		Message string `json:"message"`
	}{
		Message: "New todo created successfully",
	}

	return ctx.JSON(http.StatusOK, resp)
}

// UpdateTodo ...
func (t *TodoController) UpdateTodo(ctx echo.Context) error {
	updateTodoReq := new(requests.UpdateTodo)

	err := ctx.Bind(updateTodoReq)
	if err != nil {
		return err
	}

	todoID := ctx.Param("id")
	err = t.TodoService.UpdateTodo(todoID, updateTodoReq)
	if err != nil {
		return err
	}

	resp := struct {
		Message string `json:"message"`
	}{
		Message: "Existing todo updated successfully",
	}

	return ctx.JSON(http.StatusOK, resp)
}

// TestHandler ...
func (t *TodoController) TestHandler(ctx echo.Context) error {
	resp := struct {
		Message string `json:"message"`
	}{
		Message: "Existing todo updated successfully",
	}

	fmt.Println(resp)
	return ctx.JSON(http.StatusOK, &resp)
}
