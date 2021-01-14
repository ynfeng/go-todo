package todo

type AppContext struct {
}

func (context *AppContext) GetTodoList() *TodoList {
	return NewTodoList()
}

var Context = &AppContext{}
