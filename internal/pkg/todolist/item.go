package todolist

type Item struct {
	Name string `json:"Name"`
}

func NewItem(name string) *Item {
	return &Item{Name: name}
}
