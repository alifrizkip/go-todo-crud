package todo

import (
	"github.com/stretchr/testify/mock"
)

type TodoServiceMock struct {
	mock.Mock
}

// GetAllTodos ...
func (s *TodoServiceMock) GetAllTodos() ([]*TodoModel, error) {
	args := s.Called()
	return args.Get(0).([]*TodoModel), args.Error(1)
}

// GetATodo ...
func (s *TodoServiceMock) GetATodo(todoID string) (*TodoModel, error) {
	args := s.Called(todoID)
	return args.Get(0).(*TodoModel), args.Error(1)
}

// CreateTodo ...
func (s *TodoServiceMock) CreateTodo(cmd *NewTodoRequest) error {
	args := s.Called(cmd)
	return args.Error(0)
}

// UpdateTodo ...
func (s *TodoServiceMock) UpdateTodo(todoID string, cmd *UpdateTodoRequest) error {
	args := s.Called(todoID, cmd)
	return args.Error(0)
}

// DeleteTodo ...
func (s *TodoServiceMock) DeleteTodo(todoID string) error {
	args := s.Called(todoID)
	return args.Error(0)
}
