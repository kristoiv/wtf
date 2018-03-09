package wtf

import "time"

type ItemID string

type Item struct {
	ID      ItemID    `json:"id"`
	Title   string    `json:"title"`
	Created time.Time `json:"created"`
	Changed time.Time `json:"changed"`
	Checked bool      `json:"checked"`
}

type Items []Item

func (item Items) Len() int           { return len(item) }
func (item Items) Swap(i, j int)      { item[i], item[j] = item[j], item[i] }
func (item Items) Less(i, j int) bool { return item[i].Created.Before(item[j].Created) }

type Client interface {
	Open() error
	Close() error
	TodoListService() TodoListService
}

type TodoListService interface {
	Add(title string) (*Item, error)
	SetChecked(id ItemID, checked bool) error
	Remove(id ItemID) error
	Items() ([]Item, error)
}
