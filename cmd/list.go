package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ynfeng/todo/internal/pkg/appcontext"
	"github.com/ynfeng/todo/internal/pkg/console"
)

type ListCmd struct {
	Base
	all bool
}

func (listCmd *ListCmd) Run(cmd *cobra.Command, args []string) {
	todoList := appcontext.GetTodoList(appcontext.WorkSpace)
	console.PrintItems(todoList.All())
}

func NewListCommand() *ListCmd {
	listCmd := &ListCmd{}
	var cobraCmd = cobra.Command{
		Use:   "list",
		Short: "Print unfinished todolist list",
		Run:   listCmd.Run,
	}
	cobraCmd.Flags().BoolVarP(&listCmd.all, "all", "a", false, "Print all todolist list")
	listCmd.cmd = cobraCmd
	return listCmd
}
