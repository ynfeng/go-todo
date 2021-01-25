package cmd

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/ynfeng/todo/internal/pkg/appcontext"
	"github.com/ynfeng/todo/internal/pkg/console"
	"os"
	"testing"
)

func TestListItem(t *testing.T) {
	appcontext.SwitchWorkSpace(uuid.NewV4().String())
	addItem("foo")
	addItem("bar")

	writer := &console.TestableWriter{}
	console.SetOut(writer)
	os.Args = []string{"", "list"}

	builder := NewCommandBuilder()
	builder.RootCommand(NewRootCmd())
	builder.AddCommands(
		NewListCommand(),
	)
	builder.Build().Start()

	assert.Equal(t, "1. foo\n2. bar\n\n", string(writer.Data))
}

func addItem(itemName string) {
	os.Args = []string{"", "add", itemName}

	builder := NewCommandBuilder()
	builder.RootCommand(NewRootCmd())
	builder.AddCommands(
		NewAddCmd(),
	)
	builder.Build().Start()
}
