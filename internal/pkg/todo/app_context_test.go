package todo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTodoList(t *testing.T) {
	todoList := Context.GetTodoList()

	assert.NotNil(t, todoList)
}
