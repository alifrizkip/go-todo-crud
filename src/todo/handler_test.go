package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func Test_FindAllTodos(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/todos")
	todoService := &TodoServiceMock{}
	h := NewTodoHandler(todoService)

	todos := []*TodoModel{
		&TodoModel{TodoID: "aaa", Title: "Wake up", Detail: "Bangun tidur"},
		&TodoModel{TodoID: "bbb", Title: "Take a bath", Detail: "Mandi pagi"},
		&TodoModel{TodoID: "ccc", Title: "Go to work", Detail: "Berangkat kerja"},
	}
	todoService.On("GetAllTodos").Return(todos, nil)

	todosJSON, _ := json.Marshal(todos)
	expectedJSON := fmt.Sprintf("%s\n", todosJSON)

	// Assertions
	if assert.NoError(t, h.FindAllTodos(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedJSON, rec.Body.String())
	}
}

func Test_FindAllTodos_Error(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/todos")
	todoService := &TodoServiceMock{}
	h := NewTodoHandler(todoService)

	todos := []*TodoModel{}
	todoService.On("GetAllTodos").Return(todos, errors.New("Return error"))

	// Assertions
	assert.Error(t, h.FindAllTodos(c))
}

func Test_FindTodoByID(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("aaa")
	todoService := &TodoServiceMock{}
	h := NewTodoHandler(todoService)

	todo := &TodoModel{
		TodoID: "aaa", Title: "Wake up", Detail: "Bangun tidur",
	}
	todoService.On("GetATodo", "aaa").Return(todo, nil)

	todoJSON := `{"todo_id":"aaa","title":"Wake up","detail":"Bangun tidur","is_done":false}`
	expectedJSON := fmt.Sprintf("%s\n", todoJSON)

	// Assertions
	if assert.NoError(t, h.FindTodoByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedJSON, rec.Body.String())
	}
}

func Test_FindTodoByID_Error(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("aaa")
	todoService := &TodoServiceMock{}
	h := NewTodoHandler(todoService)

	todoService.On("GetATodo", "aaa").Return(&TodoModel{}, errors.New("Return error"))

	// Assertions
	assert.Error(t, h.FindTodoByID(c))
}

// Test_SaveTodo
func Test_SaveTodo_Success(t *testing.T) {
	todoJSON := `{"title":"Sleep","detail":"Tidur Malam"}`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	todoService := &TodoServiceMock{}
	h := NewTodoHandler(todoService)

	var newTodoReq NewTodoRequest
	_ = json.Unmarshal([]byte(todoJSON), &newTodoReq)
	todoService.On("CreateTodo", &newTodoReq).Return(nil)

	respJSON := `{"message":"New todo created successfully"}`
	expectedJSON := fmt.Sprintf("%s\n", respJSON)

	// Assertions
	if assert.NoError(t, h.SaveTodo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedJSON, rec.Body.String())
	}
}

// Test_SaveTodo_TodoService_CreateTodo_Error
func Test_SaveTodo_TodoService_CreateTodo_Error(t *testing.T) {
	todoJSON := `{"title":"Sleep","detail":"Tidur Malam"}`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	todoService := &TodoServiceMock{}
	h := NewTodoHandler(todoService)

	var newTodoReq NewTodoRequest
	_ = json.Unmarshal([]byte(todoJSON), &newTodoReq)
	todoService.On("CreateTodo", &newTodoReq).Return(errors.New("Return error"))

	// Assertions
	assert.Error(t, h.SaveTodo(c))
}

// Test_SaveTodo_ContextBind_Error
func Test_SaveTodo_ContextBind_Error(t *testing.T) {
	todoJSON := `{"title":"Sleep","detail":"Tidur Malam"`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	todoService := &TodoServiceMock{}
	h := NewTodoHandler(todoService)

	// Assertions
	assert.Error(t, h.SaveTodo(c))
}

func Test_UpdateTodo_ContextBind_Error(t *testing.T) {
	todoJSON := `{"title":"Sleep","detail":"Tidur Malam","is_done":true`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/api/todos/aaa")
	c.SetParamNames("id")
	c.SetParamValues("aaa")

	todoService := &TodoServiceMock{}
	h := NewTodoHandler(todoService)

	// Assertions
	assert.Error(t, h.UpdateTodo(c))
}

func Test_UpdateTodo_TodoService_UpdateTodo_Error(t *testing.T) {
	todoJSON := `{"title":"Sleep","detail":"Tidur Malam","is_done":true}`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/api/todos/aaa")
	c.SetParamNames("id")
	c.SetParamValues("aaa")

	todoService := &TodoServiceMock{}
	h := NewTodoHandler(todoService)

	var updateTodoReq UpdateTodoRequest
	_ = json.Unmarshal([]byte(todoJSON), &updateTodoReq)
	todoService.On("UpdateTodo", "aaa", &updateTodoReq).Return(errors.New("Return error"))

	// Assertions
	assert.Error(t, h.UpdateTodo(c))
}

func Test_UpdateTodo_Success(t *testing.T) {
	todoJSON := `{"title":"Sleep","detail":"Tidur Malam","is_done":true}`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/api/todos/aaa")
	c.SetParamNames("id")
	c.SetParamValues("aaa")

	todoService := &TodoServiceMock{}
	h := NewTodoHandler(todoService)

	var updateTodoReq UpdateTodoRequest
	_ = json.Unmarshal([]byte(todoJSON), &updateTodoReq)
	todoService.On("UpdateTodo", "aaa", &updateTodoReq).Return(nil)

	// Assertions
	assert.NoError(t, h.UpdateTodo(c))
}
