package bolt

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/kristoiv/wtf"
	"github.com/kristoiv/wtf/bolt/internal"
)

var _ wtf.TodoListService = &TodoListService{}

type TodoListService struct {
	client *Client
}

func (s *TodoListService) Add(title string) (*wtf.Item, error) {
	if title == "" {
		return nil, wtf.ErrItemTitleRequired
	}

	tx, err := s.client.db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	item := &wtf.Item{
		ID:      wtf.ItemID(uuid.New()),
		Title:   title,
		Created: s.client.Now(),
	}

	b := tx.Bucket([]byte("items"))
	if data, err := internal.MarshalItem(item); err != nil {
		return nil, err
	} else if err := b.Put([]byte(item.ID), data); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *TodoListService) Items() []wtf.Item {
	return nil
}

func (s *TodoListService) SetChecked(id wtf.ItemID, checked bool) error {
	return nil
}

func (s *TodoListService) Remove(item *wtf.Item) error {
	return nil
}
