package todo

import (
	"github.com/stretchr/testify/mock"
)

// TodoRepositoryMock ...
type TodoRepositoryMock struct {
	mock.Mock
}

// FindAll ...
func (r *TodoRepositoryMock) FindAll() ([]*TodoModel, error) {
	args := r.Called()
	return args.Get(0).([]*TodoModel), args.Error(1)
}

// FindByID ...
func (r *TodoRepositoryMock) FindByID(todoID string) (*TodoModel, error) {
	args := r.Called(todoID)
	return args.Get(0).(*TodoModel), args.Error(1)
}

// Save ...
func (r *TodoRepositoryMock) Save(todo *TodoModel) error {
	args := r.Called(todo)
	return args.Error(0)
}

// Update ...
func (r *TodoRepositoryMock) Update(todoID string, todo *TodoModel) error {
	args := r.Called(todoID, todo)
	return args.Error(0)
}

// Delete ...
func (r *TodoRepositoryMock) Delete(todoID string) error {
	args := r.Called(todoID)
	return args.Error(0)
}
