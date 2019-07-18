package services

import (
	"go-crud/models"
	"go-crud/repositories"
	"go-crud/requests"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func isInstanceOfModelsTodo(iface interface{}) bool {
	switch iface.(type) {
	case *models.Todo:
		return true
	default:
		return false
	}
}

func TestTodoService_GetAllTodos(t *testing.T) {
	todoRepo := &repositories.TodoRepoMock{}
	service := new(TodoService).New(todoRepo)

	todos := []*models.Todo{
		&models.Todo{TodoID: "aaa", Title: "Wake up", Detail: "Bangun tidur"},
		&models.Todo{TodoID: "bbb", Title: "Take a bath", Detail: "Mandi pagi"},
		&models.Todo{TodoID: "ccc", Title: "Go to work", Detail: "Berangkat kerja"},
	}

	todoRepo.On("FindAll").Return(todos, nil)

	data, err := service.GetAllTodos()
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, 3, len(todos), "Todos data must have 3 item")
	assert.True(t, isInstanceOfModelsTodo(data[0]), "Todos datas item must instance of models.Todo")
}

func TestTodoService_GetATodos(t *testing.T) {
	todoRepo := &repositories.TodoRepoMock{}
	service := new(TodoService).New(todoRepo)

	todo := &models.Todo{
		TodoID: "aaa", Title: "Wake up", Detail: "Bangun tidur",
	}

	todoRepo.On("FindByID", mock.MatchedBy(func(todoID string) bool {
		assert.Equal(t, todoID, todo.TodoID)
		return true
	})).Return(todo, nil)

	data, err := service.GetATodo("aaa")
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, "aaa", data.TodoID)
	assert.Equal(t, "Wake up", data.Title)
	assert.Equal(t, "Bangun tidur", data.Detail)
	assert.False(t, data.IsDone)
}

func TestTodoService_CreateTodo(t *testing.T) {
	todoRepo := &repositories.TodoRepoMock{}
	service := new(TodoService).New(todoRepo)

	newTodo := &requests.NewTodo{
		Title: "Wake up", Detail: "Bangun tidur",
	}

	todoRepo.On("Save", mock.Anything).Return(nil)

	err := service.CreateTodo(newTodo)
	assert.NoError(t, err)
}

func TestTodoService_UpdateTodo(t *testing.T) {
	todoRepo := &repositories.TodoRepoMock{}
	service := new(TodoService).New(todoRepo)

	updateTodo := &requests.UpdateTodo{
		Title: "Sleep", Detail: "Tidur malam", IsDone: true,
	}
	updateTodoID := "aaa"

	todoRepo.On("Update", mock.MatchedBy(func(todoID string) bool {
		assert.Equal(t, updateTodoID, todoID)
		return true
	})).Return(nil)
	err := service.UpdateTodo(updateTodoID, updateTodo)
	assert.NoError(t, err)
}

func TestTodoService_DeleteTodo(t *testing.T) {
	todoRepo := &repositories.TodoRepoMock{}
	service := new(TodoService).New(todoRepo)

	deleteTodoID := "aaa"

	todoRepo.On("Delete", mock.MatchedBy(func(todoID string) bool {
		assert.Equal(t, todoID, deleteTodoID)
		return true
	})).Return(nil)
	err := service.DeleteTodo(deleteTodoID)
	assert.NoError(t, err)
}
