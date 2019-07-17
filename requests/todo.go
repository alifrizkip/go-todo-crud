package requests

// NewTodo ...
type NewTodo struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// UpdateTodo ...
type UpdateTodo struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	IsDone bool   `json:"is_done"`
}
