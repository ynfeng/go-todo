package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ynfeng/todo/internal/pkg/console"
	"github.com/ynfeng/todo/internal/pkg/todo"
)

type AddCmd struct {
	Base
	args [] string
}

func (addCmd *AddCmd) Run(cmd *cobra.Command, args []string) {
	name := args[0]
	todoList := todo.Context.GetTodoList()
	todoList.Add(*todo.NewItem(name))

	console.PrintItems(todoList.All())
	console.Printf("%s added.\n", name)
}

func NewAddCmd() *AddCmd {
	cmd := &AddCmd{}
	var cobraCmd = cobra.Command{
		Use:   "add",
		Short: "Add todo item.",
		Run:   cmd.Run,
	}
	cobraCmd.SetUsageTemplate("Usage: \n   add <item name>\n")
	cmd.cmd = cobraCmd
	return cmd
}
