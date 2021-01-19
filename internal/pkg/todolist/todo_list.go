package todolist

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

func NewTodoList(storage Storage) *TodoList {
	todoList := &TodoList{storage: storage}
	todoList.items = storage.All()
	return todoList
}
