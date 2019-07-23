package todo

import (
	"database/sql"

	"github.com/labstack/echo"
)

type TodoServer struct {
	todoService ITodoService
	handler     *TodoHandler
}

// NewTodoServer ...
func NewTodoServer(db *sql.DB) *TodoServer {
	todoServer := &TodoServer{}
	todoRepo := NewTodoRepository(db)
	todoServer.todoService = NewTodoService(todoRepo)
	todoServer.handler = NewTodoHandler(todoServer.todoService)

	return todoServer
}

// Mount ...
func (s *TodoServer) Mount(g *echo.Group) {
	g.GET("", s.handler.FindAllTodos)
	g.GET("/:id", s.handler.FindTodoByID)
	g.POST("", s.handler.SaveTodo)
	g.PUT("/:id", s.handler.UpdateTodo)
	g.DELETE("/:id", s.handler.DeleteTodo)
}
