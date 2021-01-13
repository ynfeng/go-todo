package todo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	todoList := NewTodoList()

	todoList.Add(*NewItem("foo"))
	item := todoList.Get(1)

	assert.Equal(t, *item, *NewItem("foo"))
}

func TestListAll(t *testing.T) {
	todoList := NewTodoList()
	todoList.Add(*NewItem("foo"))
	todoList.Add(*NewItem("bar"))

	all := todoList.All()
	
	excepted := []Item{*NewItem("foo"), *NewItem("bar")}
	assert.Equal(t, excepted, all)
}
