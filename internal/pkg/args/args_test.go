package args

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCommand(t *testing.T) {
	var args *Args
	var cmd *string

	args = NewArgs([]string{"add", "foo"})
	cmd, _ = args.Cmd()
	assert.Equal(t, "add", *cmd)

	args = NewArgs([]string{"list", "foo"})
	cmd, _ = args.Cmd()
	assert.Equal(t, "list", *cmd)
}

func TestGetNotExistsCommand(t *testing.T) {
	args := NewArgs([]string{})

	_, error := args.Cmd()

	assert.IsType(t, new(ArgLineError), error)
	assert.Equal(t, usage, error.Error())
}
