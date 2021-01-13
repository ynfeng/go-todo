package todo

type Item struct {
	name string
}

func NewItem(name string) *Item {
	return &Item{name: name}
}
