package todo

import (
	"net/http"

	"github.com/labstack/echo"
)

type TodoHandler struct {
	todoService ITodoService
}

// NewTodoHandler ...
func NewTodoHandler(service ITodoService) *TodoHandler {
	return &TodoHandler{todoService: service}
}

// FindAllTodos ...
func (t *TodoHandler) FindAllTodos(ctx echo.Context) error {
	data, err := t.todoService.GetAllTodos()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// FindTodoByID ...
func (t *TodoHandler) FindTodoByID(ctx echo.Context) error {
	todoID := ctx.Param("id")
	data, err := t.todoService.GetATodo(todoID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// SaveTodo ...
func (t *TodoHandler) SaveTodo(ctx echo.Context) error {
	newTodoReq := new(NewTodoRequest)

	err := ctx.Bind(newTodoReq)
	if err != nil {
		return err
	}

	err = t.todoService.CreateTodo(newTodoReq)
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
func (t *TodoHandler) UpdateTodo(ctx echo.Context) error {
	updateTodoReq := new(UpdateTodoRequest)

	err := ctx.Bind(updateTodoReq)
	if err != nil {
		return err
	}

	todoID := ctx.Param("id")
	err = t.todoService.UpdateTodo(todoID, updateTodoReq)
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

// DeleteTodo ...
func (t *TodoHandler) DeleteTodo(ctx echo.Context) error {
	todoID := ctx.Param("id")
	err := t.todoService.DeleteTodo(todoID)
	if err != nil {
		return err
	}

	resp := struct {
		Message string `json:"message"`
	}{
		Message: "Todo deleted successfully",
	}

	return ctx.JSON(http.StatusOK, resp)
}
