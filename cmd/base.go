package cmd

import "github.com/spf13/cobra"

type Base struct {
	cmd cobra.Command
}

func (base *Base) CobraCommand() *cobra.Command {
	return &base.cmd
}
