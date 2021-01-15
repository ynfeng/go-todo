package storage

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestAppendItem(t *testing.T) {
	randomFile := uuid.NewV4()
	tempDir, _ := ioutil.TempDir("", "todolist-")
	storage := NewFileStorage(tempDir, randomFile.String())

	var err = storage.Append("test1")
	assert.Nil(t, err)

	err = storage.Append("test2")
	assert.Nil(t, err)

	items, err := storage.All(func() interface{} {
		var str string
		return str
	})
	assert.Nil(t, err)
	assert.Equal(t, items[0], "test1")
	assert.Equal(t, items[1], "test2")
}
