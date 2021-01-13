package main

import (
	"github.com/ynfeng/todo/cmd"
)

func init() {
	builder := cmd.NewCommandBuilder()
	builder.RootCommand(cmd.NewRootCmd())
	builder.AddCommands(cmd.NewListCommand())
	builder.Build().Start()
}

func main() {
}
