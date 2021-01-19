package todolist

type Storage interface {
	Append(item Item) error
	All() ([]Item, error)
	replaceAll(items ...Item) error
}
