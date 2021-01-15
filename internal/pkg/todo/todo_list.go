package todo

type TodoList struct {
	items []Item
}

func (list *TodoList) Add(item Item) {
	list.items = append(list.items, item)
}

func (list *TodoList) Get(idx int) *Item {
	return &list.items[idx-1]
}

func (list *TodoList) All() []Item {
	return list.items
}

func NewTodoList() *TodoList {
	return &TodoList{}
}
