package cmd

import (
	"github.com/spf13/cobra"
)

type Command interface {
	Run(command *cobra.Command, args []string)
	CobraCommand() *cobra.Command
}

type RootCommand interface {
	AddCommand(command Command)
	Start()
}

type CommandBuilder struct {
	commands    []Command
	rootCommand RootCommand
}

func (builder *CommandBuilder) AddCommands(commands ...Command) *CommandBuilder {
	builder.commands = append(builder.commands, commands...)
	return builder
}

func (builder *CommandBuilder) RootCommand(command RootCommand) *CommandBuilder {
	builder.rootCommand = command
	return builder
}

func (builder *CommandBuilder) Build() RootCommand {
	for _, cmd := range builder.commands {
		builder.rootCommand.AddCommand(cmd)
	}
	cobra.OnInitialize()
	return builder.rootCommand
}

func NewCommandBuilder() *CommandBuilder {
	return &CommandBuilder{}
}
