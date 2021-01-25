package todolist

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestAdd(t *testing.T) {
	randomFile := uuid.NewV4().String()
	tempDir, _ := ioutil.TempDir("", "todolist-")
	fileStorage := NewFileStorage(tempDir, randomFile)
	todoList := NewTodoList(fileStorage)

	todoList.Add(*NewItem("foo"))
	item := todoList.Get(1)

	assert.Equal(t, *item, *NewItem("foo"))
}

func TestListAll(t *testing.T) {
	randomFile := uuid.NewV4().String()
	tempDir, _ := ioutil.TempDir("", "todolist-")
	fileStorage := NewFileStorage(tempDir, randomFile)
	todoList := NewTodoList(fileStorage)

	todoList.Add(*NewItem("foo"))
	todoList.Add(*NewItem("bar"))

	all := todoList.All()

	excepted := []Item{*NewItem("foo"), *NewItem("bar")}
	assert.Equal(t, excepted, all)
}

func TestPersistence(t *testing.T) {
	randomFile := uuid.NewV4().String()
	tempDir, _ := ioutil.TempDir("", "todolist-")
	fileStorage := NewFileStorage(tempDir, randomFile)
	var todoList = NewTodoList(fileStorage)

	todoList.Add(*NewItem("foo"))
	todoList.Add(*NewItem("bar"))

	todoList = NewTodoList(fileStorage)
	all := todoList.All()
	excepted := []Item{*NewItem("foo"), *NewItem("bar")}
	assert.Equal(t, excepted, all)
}

func TestItem_Done(t *testing.T) {
	randomFile := uuid.NewV4().String()
	tempDir, _ := ioutil.TempDir("", "todolist-")
	fileStorage := NewFileStorage(tempDir, randomFile)
	var todoList = NewTodoList(fileStorage)

	todoList.Add(*NewItem("foo"))
	todoList.Add(*NewItem("bar"))
	todoList.Done(2)

	all := todoList.All()
	doneItem := NewItem("bar")
	doneItem.Done()
	excepted := []Item{*NewItem("foo"), *doneItem}
	assert.Equal(t, excepted, all)
}

func TestOutOfItemIndexRange(t *testing.T) {
	randomFile := uuid.NewV4().String()
	tempDir, _ := ioutil.TempDir("", "todolist-")
	fileStorage := NewFileStorage(tempDir, randomFile)
	var todoList = NewTodoList(fileStorage)

	err := todoList.Done(22)

	assert.Equal(t, err.Error(), "out of item index range")
}
