package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type RootCmd struct {
	cmd *cobra.Command
}

func (rootCmd *RootCmd) Start() {
	if err := rootCmd.cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (rootCmd *RootCmd) AddCommand(command Command) {
	rootCmd.cmd.AddCommand(command.CobraCommand())
}

func NewRootCmd() *RootCmd {
	command := &cobra.Command{}
	command.Version = "v0.1-alpha"
	return &RootCmd{command}
}
