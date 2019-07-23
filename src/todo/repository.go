package todo

import (
	"database/sql"
)

// ITodoRepository ...
type ITodoRepository interface {
	FindAll() ([]*TodoModel, error)
	FindByID(todoID string) (*TodoModel, error)
	Save(todo *TodoModel) error
	Update(todoID string, todo *TodoModel) error
	Delete(todoID string) error
}

// TodoRepository ...
type TodoRepository struct {
	db *sql.DB
}

// NewTodoRepository ...
func NewTodoRepository(conn *sql.DB) ITodoRepository {
	return &TodoRepository{db: conn}
}

// FindAll ...
func (r *TodoRepository) FindAll() ([]*TodoModel, error) {
	rows, err := r.db.
		Query("select todo_id, title, detail, is_done from todos")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*TodoModel{}

	for rows.Next() {
		todo := &TodoModel{}

		if err := rows.Scan(&todo.TodoID, &todo.Title, &todo.Detail, &todo.IsDone); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

// FindByID ...
func (r *TodoRepository) FindByID(todoID string) (*TodoModel, error) {
	var todo = &TodoModel{}
	err := r.db.
		QueryRow("select todo_id, title, detail, is_done from todos where todo_id = ?", todoID).
		Scan(&todo.TodoID, &todo.Title, &todo.Detail, &todo.IsDone)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

// Save ...
func (r *TodoRepository) Save(todo *TodoModel) error {
	_, err := r.db.Exec("insert into todos values (?, ?, ?, ?)", todo.TodoID, todo.Title, todo.Detail, todo.IsDone)
	if err != nil {
		return err
	}

	return nil
}

// Update ...
func (r *TodoRepository) Update(todoID string, todo *TodoModel) error {
	_, err := r.db.Exec(`update todos
		set title = ?, detail = ?, is_done = ?
		where todo_id = ?`,
		todo.Title, todo.Detail, todo.IsDone, todoID)
	if err != nil {
		return err
	}

	return nil
}

// Delete ...
func (r *TodoRepository) Delete(todoID string) error {
	_, err := r.db.Exec("delete from todos where todo_id = ?", todoID)
	if err != nil {
		return err
	}

	return nil
}
