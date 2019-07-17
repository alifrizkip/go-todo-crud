package services

import (
	"go-crud/models"
	"go-crud/repositories"
	"go-crud/requests"

	uuid "github.com/satori/go.uuid"
)

// ITodoService ...
type ITodoService interface {
	CreateTodo(cmd *requests.NewTodo) error
	UpdateTodo(todoID string, cmd *requests.UpdateTodo) error
	DeleteTodo(todoID string) error
	GetAllTodos() ([]*models.Todo, error)
	GetATodo(todoID string) (*models.Todo, error)
}

// TodoService ...
type TodoService struct {
	todoRepo repositories.ITodoRepo
}

// New ...
func (*TodoService) New(repo repositories.ITodoRepo) ITodoService {
	service := &TodoService{todoRepo: repo}

	return service
}

// GetAllTodos ...
func (s *TodoService) GetAllTodos() ([]*models.Todo, error) {
	data, err := s.todoRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetATodo ...
func (s *TodoService) GetATodo(todoID string) (*models.Todo, error) {
	todo, err := s.todoRepo.FindByID(todoID)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// CreateTodo ...
func (s *TodoService) CreateTodo(cmd *requests.NewTodo) error {
	todoID := uuid.NewV4().String()

	todo := models.Todo{
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
func (s *TodoService) UpdateTodo(todoID string, cmd *requests.UpdateTodo) error {
	todo := models.Todo{
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
