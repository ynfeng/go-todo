package appcontext

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTodoList(t *testing.T) {
	todoList := GetTodoList()

	assert.NotNil(t, todoList)
}
