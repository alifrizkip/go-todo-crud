package repositories

import (
	"go-crud/models"

	"github.com/stretchr/testify/mock"
)

// TodoRepoMock ...
type TodoRepoMock struct {
	mock.Mock
}

// FindAll ...
func (r *TodoRepoMock) FindAll() ([]*models.Todo, error) {
	args := r.Called()
	return args.Get(0).([]*models.Todo), args.Error(1)
}

// FindByID ...
func (r *TodoRepoMock) FindByID(todoID string) (*models.Todo, error) {
	args := r.Called(todoID)
	return args.Get(0).(*models.Todo), args.Error(1)
}

// Save ...
func (r *TodoRepoMock) Save(todo *models.Todo) error {
	args := r.Called(todo)
	return args.Error(0)
}

// Update ...
func (r *TodoRepoMock) Update(todoID string, todo *models.Todo) error {
	args := r.Called(todoID)
	return args.Error(0)
}

// Delete ...
func (r *TodoRepoMock) Delete(todoID string) error {
	args := r.Called(todoID)
	return args.Error(0)
}
