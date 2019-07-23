package todo

// NewTodoRequest ...
type NewTodoRequest struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// UpdateTodoRequest ...
type UpdateTodoRequest struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	IsDone bool   `json:"is_done"`
}
