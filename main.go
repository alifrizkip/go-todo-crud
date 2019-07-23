package main

import (
	"database/sql"
	"fmt"

	todoModule "go-crud/src/todo"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("Starting server...")

	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	todoServer := todoModule.NewTodoServer(db)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	API := e.Group("api")

	todoGroup := API.Group("/todos")
	todoServer.Mount(todoGroup)

	e.Start(":9000")
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_todos")
	if err != nil {
		return nil, err
	}

	return db, nil
}
