package storage

type ItemFactory func() interface{}

type Storage interface {
	Append(item interface{}) error
	All(factory ItemFactory) ([]interface{}, error)
}
