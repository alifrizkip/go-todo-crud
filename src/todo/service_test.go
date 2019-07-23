package todo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func isInstanceOfModelsTodo(iface interface{}) bool {
	switch iface.(type) {
	case *TodoModel:
		return true
	default:
		return false
	}
}

func TestTodoService_GetAllTodos(t *testing.T) {
	todoRepo := &TodoRepositoryMock{}
	service := NewTodoService(todoRepo)

	todos := []*TodoModel{
		&TodoModel{TodoID: "aaa", Title: "Wake up", Detail: "Bangun tidur"},
		&TodoModel{TodoID: "bbb", Title: "Take a bath", Detail: "Mandi pagi"},
		&TodoModel{TodoID: "ccc", Title: "Go to work", Detail: "Berangkat kerja"},
	}

	todoRepo.On("FindAll").Return(todos, nil)

	data, err := service.GetAllTodos()
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, 3, len(todos), "Todos data must have 3 item")
	assert.True(t, isInstanceOfModelsTodo(data[0]), "Todos datas item must instance of TodoModel")
}

func TestTodoService_GetAllTodos_Error(t *testing.T) {
	todoRepo := &TodoRepositoryMock{}
	service := NewTodoService(todoRepo)

	todos := []*TodoModel{}
	todoRepo.On("FindAll").Return(todos, errors.New("Return error"))

	data, err := service.GetAllTodos()
	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestTodoService_GetATodo(t *testing.T) {
	todoRepo := &TodoRepositoryMock{}
	service := NewTodoService(todoRepo)

	todo := &TodoModel{
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

func TestTodoService_GetATodo_Error(t *testing.T) {
	todoRepo := &TodoRepositoryMock{}
	service := NewTodoService(todoRepo)

	todoRepo.On("FindByID", "aaa").Return(&TodoModel{}, errors.New("Return error"))

	data, err := service.GetATodo("aaa")
	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestTodoService_CreateTodo(t *testing.T) {
	todoRepo := &TodoRepositoryMock{}
	service := NewTodoService(todoRepo)

	newTodo := &NewTodoRequest{
		Title: "Wake up", Detail: "Bangun tidur",
	}

	todoRepo.On("Save", mock.Anything).Return(nil)

	err := service.CreateTodo(newTodo)
	assert.NoError(t, err)
}

func TestTodoService_CreateTodo_Error(t *testing.T) {
	todoRepo := &TodoRepositoryMock{}
	service := NewTodoService(todoRepo)

	newTodo := &NewTodoRequest{
		Title: "Wake up", Detail: "Bangun tidur",
	}

	todoRepo.On("Save", mock.Anything).Return(errors.New("Return error"))

	err := service.CreateTodo(newTodo)
	assert.Error(t, err)
}

func TestTodoService_UpdateTodo(t *testing.T) {
	todoRepo := &TodoRepositoryMock{}
	service := NewTodoService(todoRepo)

	updateTodoReq := &UpdateTodoRequest{
		Title: "Sleep", Detail: "Tidur malam", IsDone: true,
	}
	updateTodoID := "aaa"

	updateTodo := &TodoModel{
		Title:  updateTodoReq.Title,
		Detail: updateTodoReq.Detail,
		IsDone: updateTodoReq.IsDone,
	}

	todoRepo.On("Update", updateTodoID, updateTodo).Return(nil)
	err := service.UpdateTodo(updateTodoID, updateTodoReq)
	assert.NoError(t, err)
}

func TestTodoService_UpdateTodo_Error(t *testing.T) {
	todoRepo := &TodoRepositoryMock{}
	service := NewTodoService(todoRepo)

	updateTodoReq := &UpdateTodoRequest{
		Title: "Sleep", Detail: "Tidur malam", IsDone: true,
	}
	updateTodoID := "aaa"

	updateTodo := &TodoModel{
		Title:  updateTodoReq.Title,
		Detail: updateTodoReq.Detail,
		IsDone: updateTodoReq.IsDone,
	}

	todoRepo.On("Update", updateTodoID, updateTodo).Return(errors.New("Return error"))
	err := service.UpdateTodo(updateTodoID, updateTodoReq)
	assert.Error(t, err)
}

func TestTodoService_DeleteTodo(t *testing.T) {
	todoRepo := &TodoRepositoryMock{}
	service := NewTodoService(todoRepo)

	deleteTodoID := "aaa"

	todoRepo.On("Delete", mock.MatchedBy(func(todoID string) bool {
		assert.Equal(t, todoID, deleteTodoID)
		return true
	})).Return(nil)
	err := service.DeleteTodo(deleteTodoID)
	assert.NoError(t, err)
}

func TestTodoService_DeleteTodo_Error(t *testing.T) {
	todoRepo := &TodoRepositoryMock{}
	service := NewTodoService(todoRepo)

	todoRepo.On("Delete", "aaa").Return(errors.New("Return error"))
	err := service.DeleteTodo("aaa")
	assert.Error(t, err)
}
