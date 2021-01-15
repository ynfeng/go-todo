package appcontext

import "github.com/ynfeng/todo/internal/pkg/todo"

func GetTodoList() *todo.TodoList {
	return todo.NewTodoList()
}
