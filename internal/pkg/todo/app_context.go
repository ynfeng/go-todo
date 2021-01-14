package todo

type AppContext struct {
	todoList TodoList
}

func (context *AppContext) GetTodoList() *TodoList {
	return &context.todoList
}

var Context = &AppContext{*NewTodoList()}
