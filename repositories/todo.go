package repositories

import (
	"database/sql"
	"go-crud/models"
)

// ITodoRepo ...
type ITodoRepo interface {
	FindAll() ([]*models.Todo, error)
	FindByID(todoID string) (*models.Todo, error)
	Save(todo *models.Todo) error
	Update(todoID string, todo *models.Todo) error
	Delete(todoID string) error
}

// TodoRepo ...
type TodoRepo struct {
	db *sql.DB
}

// New ...
func (*TodoRepo) New(conn *sql.DB) ITodoRepo {
	return &TodoRepo{db: conn}
}

// FindAll ...
func (r *TodoRepo) FindAll() ([]*models.Todo, error) {
	rows, err := r.db.
		Query("select todo_id, title, detail, is_done from todos")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*models.Todo{}

	for rows.Next() {
		todo := &models.Todo{}

		if err := rows.Scan(&todo.TodoID, &todo.Title, &todo.Detail, &todo.IsDone); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

// FindByID ...
func (r *TodoRepo) FindByID(todoID string) (*models.Todo, error) {
	var todo = &models.Todo{}
	err := r.db.
		QueryRow("select todo_id, title, detail, is_done from todos where todo_id = ?", todoID).
		Scan(&todo.TodoID, &todo.Title, &todo.Detail, &todo.IsDone)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

// Save ...
func (r *TodoRepo) Save(todo *models.Todo) error {
	return nil
}

// Update ...
func (r *TodoRepo) Update(todoID string, todo *models.Todo) error {
	return nil
}

// Delete ...
func (r *TodoRepo) Delete(todoID string) error {
	return nil
}
