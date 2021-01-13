package cmd

import (
	"github.com/spf13/cobra"
)

type ListCmd struct {
	cmd cobra.Command
	all bool
}

func (listCmd *ListCmd) Run(cmd *cobra.Command, args []string) {
}

func (listCmd *ListCmd) CobraCommand() *cobra.Command {
	return &listCmd.cmd
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
