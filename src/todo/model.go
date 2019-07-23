package todo

// TodoModel model
type TodoModel struct {
	TodoID string `json:"todo_id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	IsDone bool   `json:"is_done"`
}
