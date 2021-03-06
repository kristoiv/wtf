package bolt

import (
	"sort"
	"time"

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
	if d, err := internal.MarshalItem(item); err != nil {
		return nil, err
	} else if err := b.Put([]byte(item.ID), d); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *TodoListService) SetChecked(id wtf.ItemID, checked bool) error {
	if id == "" {
		return wtf.ErrItemIDRequired
	}

	tx, err := s.client.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte("items"))

	item := &wtf.Item{}
	if d := b.Get([]byte(id)); d == nil {
		return wtf.ErrItemNotFound
	} else if err := internal.UnmarshalItem(d, item); err != nil {
		return err
	}

	item.Checked = checked
	item.Changed = time.Now().UTC()

	if data, err := internal.MarshalItem(item); err != nil {
		return err
	} else if err := b.Put([]byte(item.ID), data); err != nil {
		return err
	}

	return tx.Commit()
}

func (s *TodoListService) Remove(id wtf.ItemID) error {
	if id == "" {
		return wtf.ErrItemIDRequired
	}

	tx, err := s.client.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte("items"))
	if data := b.Get([]byte(id)); data == nil {
		return wtf.ErrItemNotFound
	} else if err := b.Delete([]byte(id)); err != nil {
		return err
	}

	return tx.Commit()
}

func (s *TodoListService) Items() ([]wtf.Item, error) {
	tx, err := s.client.db.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	items := wtf.Items{}
	b := tx.Bucket([]byte("items"))
	err = b.ForEach(func(_ []byte, v []byte) error {
		item := wtf.Item{}
		if err := internal.UnmarshalItem(v, &item); err != nil {
			return err
		}
		items = append(items, item)
		return nil
	})

	sort.Sort(items)
	return items, err
}
