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

type Client interface {
	TodoListService() TodoListService
}

type TodoListService interface {
	Add(title string) (*Item, error)
	Items() []Item
	SetChecked(id ItemID, checked bool) error
	Remove(item *Item) error
}
