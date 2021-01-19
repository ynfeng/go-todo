package appcontext

import (
	"github.com/mitchellh/go-homedir"
	"github.com/ynfeng/todo/internal/pkg/todolist"
)

func GetTodoList() *todolist.TodoList {
	dataDir, _ := homedir.Expand("~/.todo")
	fileStorage := todolist.NewFileStorage(dataDir, "test")
	return todolist.NewTodoList(fileStorage)
}
