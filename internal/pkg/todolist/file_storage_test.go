package todolist

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

	var err = storage.Append(*NewItem("test1"))
	assert.Nil(t, err)

	err = storage.Append(*NewItem("test2"))
	assert.Nil(t, err)

	items, err := storage.All()
	assert.Nil(t, err)
	assert.Equal(t, *NewItem("test1"), items[0])
	assert.Equal(t, *NewItem("test2"), items[1])
}

func TestReplaceAll(t *testing.T) {
	randomFile := uuid.NewV4()
	tempDir, _ := ioutil.TempDir("", "todolist-")
	storage := NewFileStorage(tempDir, randomFile.String())

	_ = storage.Append(*NewItem("test1"))
	var allItems, _ = storage.All()
	assert.Equal(t, allItems[0], *NewItem("test1"))

	var err = storage.replaceAll(*NewItem("test3"), *NewItem("test4"))
	assert.Nil(t, err)

	replacedItems, _ := storage.All()
	assert.Equal(t, *NewItem("test3"), replacedItems[0])
	assert.Equal(t, *NewItem("test4"), replacedItems[1])
}
