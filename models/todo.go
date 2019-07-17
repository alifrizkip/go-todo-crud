package models

// Todo model
type Todo struct {
	TodoID string `json:"todo_id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	IsDone bool   `json:"is_done"`
}
