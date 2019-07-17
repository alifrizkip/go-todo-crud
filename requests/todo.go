package requests

type NewTodo struct {
	Title  string
	Detail string
}

type UpdateTodo struct {
	Title  string
	Detail string
	IsDone bool
}
