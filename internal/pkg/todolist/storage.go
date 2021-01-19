package todolist

type Storage interface {
	Append(item Item)
	All() []Item
	replaceAll(items ...Item)
}
