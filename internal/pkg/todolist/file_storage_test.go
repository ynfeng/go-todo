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

	storage.Append(*NewItem("test1"))
	storage.Append(*NewItem("test2"))

	items := storage.All()
	assert.Equal(t, *NewItem("test1"), items[0])
	assert.Equal(t, *NewItem("test2"), items[1])
}

func TestReplaceAll(t *testing.T) {
	randomFile := uuid.NewV4()
	tempDir, _ := ioutil.TempDir("", "todolist-")
	storage := NewFileStorage(tempDir, randomFile.String())

	storage.Append(*NewItem("test1"))
	allItems := storage.All()
	assert.Equal(t, allItems[0], *NewItem("test1"))

	storage.replaceAll(*NewItem("test3"), *NewItem("test4"))

	replacedItems := storage.All()
	assert.Equal(t, *NewItem("test3"), replacedItems[0])
	assert.Equal(t, *NewItem("test4"), replacedItems[1])
}
