package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ynfeng/todo/internal/pkg/appcontext"
	"github.com/ynfeng/todo/internal/pkg/console"
	"github.com/ynfeng/todo/internal/pkg/todolist"
)

type AddCmd struct {
	Base
	args []string
}

func (addCmd *AddCmd) Run(cmd *cobra.Command, args []string) {
	if !hasItemName(args) {
		console.Printf("Usage: add <item name>\n")
		return
	}
	todoList := addItemToTodoList(args[0])
	console.PrintItems(todoList.All())
	console.Printf("%s added.\n", args[0])
}

func addItemToTodoList(name string) *todolist.TodoList {
	todoList := appcontext.GetTodoList()
	todoList.Add(*todolist.NewItem(name))
	return todoList
}

func hasItemName(args []string) bool {
	return !(len(args) == 0 || args[0] == "")
}

func NewAddCmd() *AddCmd {
	cmd := &AddCmd{}
	var cobraCmd = cobra.Command{
		Use:   "add",
		Short: "Add todolist item.",
		Run:   cmd.Run,
	}
	cobraCmd.SetUsageTemplate("Usage: \n   add <item name>\n")
	cmd.cmd = cobraCmd
	return cmd
}
