package todolist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItemDone(t *testing.T) {
	item := NewItem("foo")
	assert.Equal(t, item.Status.IsDone(), false)

	item.Done()
	assert.Equal(t, item.Status.IsDone(), true)
}
