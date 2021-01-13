package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type AddCmd struct {
	Base
	args [] string
}

func (addCmd *AddCmd) Run(cmd *cobra.Command, args []string) {
	fmt.Printf("%s added.\n", args[0])
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
