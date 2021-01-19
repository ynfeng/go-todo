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
	assert.Equal(t, "test1", items[0])
	assert.Equal(t, "test2", items[1])
}

func TestReplaceAll(t *testing.T) {
	randomFile := uuid.NewV4()
	tempDir, _ := ioutil.TempDir("", "todolist-")
	storage := NewFileStorage(tempDir, randomFile.String())

	storage.Append("test1")
	var allItems, _ = storage.All(func() interface{} {
		var str string
		return str
	})
	assert.Equal(t, allItems[0], "test1")

	var err = storage.replaceAll("test3", "test4")
	assert.Nil(t, err)

	replacedItems, _ := storage.All(func() interface{} {
		var str string
		return str
	})
	assert.Equal(t, "test3", replacedItems[0])
	assert.Equal(t, "test4", replacedItems[1])
}
