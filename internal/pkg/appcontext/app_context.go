package appcontext

import (
	"github.com/mitchellh/go-homedir"
	"github.com/ynfeng/todo/internal/pkg/todolist"
)

var WorkSpace = "default"

func GetTodoList(workspace string) *todolist.TodoList {
	dataDir, _ := homedir.Expand("~/.todo/")
	fileStorage := todolist.NewFileStorage(dataDir, "/"+workspace)
	return todolist.NewTodoList(fileStorage)
}

func SwitchWorkSpace(workspace string) {
	WorkSpace = workspace
}
