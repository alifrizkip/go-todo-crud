package main

import (
	"database/sql"
	"fmt"
	"go-crud/repositories"
	"go-crud/routers"
	"go-crud/services"

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

	server := echo.New()

	todoRepo := new(repositories.TodoRepo).New(db)
	todoService := new(services.TodoService).New(todoRepo)

	router := new(routers.TodoRouter)
	router.Services.TodoService = todoService

	server = router.New(server)

	server.Start(":9000")
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_todos")
	if err != nil {
		return nil, err
	}

	return db, nil
}
