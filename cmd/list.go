package cmd

import (
	"github.com/spf13/cobra"
)

type ListCmd struct {
	Base
	all bool
}

func (listCmd *ListCmd) Run(cmd *cobra.Command, args []string) {
	
}

func NewListCommand() *ListCmd {
	listCmd := &ListCmd{}
	var cobraCmd = cobra.Command{
		Use:   "list",
		Short: "Print unfinished todo list",
		Run:   listCmd.Run,
	}
	cobraCmd.Flags().BoolVarP(&listCmd.all, "all", "a", false, "Print all todo list")
	listCmd.cmd = cobraCmd
	return listCmd
}
