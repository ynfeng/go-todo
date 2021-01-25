package todolist

type Item struct {
	Name   string     `json:"Name"`
	Status ItemStatus `json:"Status"`
}

func (item *Item) Done() {
	item.Status = ItemStatus{value: DONE}
}

type ItemStatus struct {
	value string `json:"Value"`
}

func (status ItemStatus) IsDone() bool {
	return status.value == DONE
}

const (
	UNDONE = "UNDONE"
	DONE   = "DONE"
)

func NewItem(name string) *Item {
	return &Item{Name: name, Status: ItemStatus{UNDONE}}
}
