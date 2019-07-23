package todo

import (
	uuid "github.com/satori/go.uuid"
)

// ITodoService ...
type ITodoService interface {
	CreateTodo(cmd *NewTodoRequest) error
	UpdateTodo(todoID string, cmd *UpdateTodoRequest) error
	DeleteTodo(todoID string) error
	GetAllTodos() ([]*TodoModel, error)
	GetATodo(todoID string) (*TodoModel, error)
}

// TodoService ...
type TodoService struct {
	todoRepo ITodoRepository
}

// NewTodoService ...
func NewTodoService(repo ITodoRepository) ITodoService {
	return &TodoService{todoRepo: repo}
}

// GetAllTodos ...
func (s *TodoService) GetAllTodos() ([]*TodoModel, error) {
	data, err := s.todoRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetATodo ...
func (s *TodoService) GetATodo(todoID string) (*TodoModel, error) {
	todo, err := s.todoRepo.FindByID(todoID)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// CreateTodo ...
func (s *TodoService) CreateTodo(cmd *NewTodoRequest) error {
	todoID := uuid.NewV4().String()

	todo := TodoModel{
		TodoID: todoID,
		Title:  cmd.Title,
		Detail: cmd.Detail,
		IsDone: false,
	}

	err := s.todoRepo.Save(&todo)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTodo ...
func (s *TodoService) UpdateTodo(todoID string, cmd *UpdateTodoRequest) error {
	todo := TodoModel{
		Title:  cmd.Title,
		Detail: cmd.Detail,
		IsDone: cmd.IsDone,
	}

	err := s.todoRepo.Update(todoID, &todo)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTodo ...
func (s *TodoService) DeleteTodo(todoID string) error {
	err := s.todoRepo.Delete(todoID)
	if err != nil {
		return err
	}

	return nil
}
