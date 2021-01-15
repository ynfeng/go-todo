package appcontext

import "github.com/ynfeng/todo/internal/pkg/todolist"

func GetTodoList() *todolist.TodoList {
	return todolist.NewTodoList()
}
