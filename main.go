package main

import (
	"database/sql"
	"fmt"
	"net/http"

	repo "go-crud/repositories"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Starting server...")

	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	r := echo.New()

	r.GET("/todos", func(ctx echo.Context) error {
		todoRepo := new(repo.TodoRepo).New(db)
		data, _ := todoRepo.FindAll()

		return ctx.JSON(http.StatusOK, data)
	})

	r.GET("/todos/:id", func(ctx echo.Context) error {
		todoRepo := new(repo.TodoRepo).New(db)
		data, _ := todoRepo.FindByID(ctx.Param("id"))

		return ctx.JSON(http.StatusOK, data)
	})
	// r.POST("/todos", createTodoHandler)
	// r.PUT("/todos/:id", updateTodoHandler)
	// r.DELETE("/todos/:id", deleteTodoHandler)

	r.Start(":9000")
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_todos")
	if err != nil {
		return nil, err
	}

	return db, nil
}
