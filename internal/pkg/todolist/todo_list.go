package todolist

import (
	"errors"
)

type TodoList struct {
	items   []Item
	storage Storage
}

func (list *TodoList) Add(item Item) {
	list.items = append(list.items, item)
	list.storage.Append(item)
}

func (list *TodoList) Get(idx int) *Item {
	return &list.items[idx-1]
}

func (list *TodoList) All() []Item {
	return list.items
}

func (list *TodoList) Done(itemIdx int) error {
	if itemIdx-1 > len(list.items) || itemIdx == 0 {
		return errors.New("out of item index range")
	}
	list.items[itemIdx-1].Done()
	return nil
}

func NewTodoList(storage Storage) *TodoList {
	todoList := &TodoList{storage: storage}
	todoList.items = storage.All()
	return todoList
}
