package cmd

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/ynfeng/todo/internal/pkg/appcontext"
	"github.com/ynfeng/todo/internal/pkg/console"
	"os"
	"testing"
)

func TestAddItem(t *testing.T) {
	appcontext.SwitchWorkSpace(uuid.NewV4().String())
	os.Args = []string{"", "add", "foo"}
	writer := &console.TestableWriter{}
	console.SetOut(writer)

	builder := NewCommandBuilder()
	builder.RootCommand(NewRootCmd())
	builder.AddCommands(
		NewAddCmd(),
	)
	builder.Build().Start()

	assert.Equal(t, "1. foo\n\nfoo added.\n", string(writer.Data))
}

func TestNotGivenItemName(t *testing.T) {
	appcontext.SwitchWorkSpace(uuid.NewV4().String())
	os.Args = []string{"", "add", ""}
	writer := &console.TestableWriter{}
	console.SetOut(writer)

	builder := NewCommandBuilder()
	builder.RootCommand(NewRootCmd())
	builder.AddCommands(
		NewAddCmd(),
	)
	builder.Build().Start()

	assert.Equal(t, "Usage: add <item name>\n", string(writer.Data))
}
