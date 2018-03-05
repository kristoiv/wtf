package bolt

import (
	"time"

	"github.com/boltdb/bolt"
	"github.com/kristoiv/wtf"
)

const DefaultPath = "bolt.db"

type Client struct {
	Path            string
	Now             func() time.Time
	todoListService TodoListService
	db              *bolt.DB
}

func NewClient() *Client {
	c := &Client{Path: DefaultPath, Now: time.Now}
	c.todoListService.client = c
	return c
}

func (c *Client) Open() error {
	db, err := bolt.Open(c.Path, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	c.db = db

	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.CreateBucketIfNotExists([]byte("items")); err != nil {
		return err
	}

	return tx.Commit()
}

func (c *Client) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

func (c *Client) TodoListService() wtf.TodoListService { return &c.todoListService }
