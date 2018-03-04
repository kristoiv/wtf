package bolt_test

import (
	"reflect"
	"testing"

	"github.com/kristoiv/wtf"
)

func TestTodoListService_Add(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.TodoListService()

	title := "A todo item"
	item, err := s.Add(title)
	if err != nil {
		t.Fatal(err)
	}

	if item.Title != title {
		t.Fatalf("unexpected title %q != %q", item.Title, title)
	}
}

func TestTodoListService_Add_ErrItemTitleRequired(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.TodoListService()

	title := ""
	_, err := s.Add(title)
	if err != wtf.ErrItemTitleRequired {
		t.Fatalf("unexpected error type %#v", err)
	}
}

func TestTodoListService_SetChecked(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.TodoListService()

	item, _ := s.Add("A todo item")
	if err := s.SetChecked(item.ID, true); err != nil {
		t.Fatal(err)
	}
}

func TestTodoListService_SetChecked_ErrItemIDRequired(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.TodoListService()

	if err := s.SetChecked("", true); err != wtf.ErrItemIDRequired {
		t.Fatalf("unexpected error type %#v", err)
	}
}

func TestTodoListService_Remove(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.TodoListService()

	item, _ := s.Add("A todo item")
	if err := s.Remove(item.ID); err != nil {
		t.Fatal(err)
	}
}

func TestTodoListService_Remove_ErrItemIDRequired(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.TodoListService()

	if err := s.Remove(""); err != wtf.ErrItemIDRequired {
		t.Fatalf("unexpected error type %#v", err)
	}
}

func TestTodoListService_List(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.TodoListService()

	item1, _ := s.Add("Todo item 1")
	item2, _ := s.Add("Todo item 2")

	items, err := s.Items()
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(*item1, items[0]) {
		t.Fatalf("unexpected item in list: %#v != %#v", items[0], item1)
	} else if !reflect.DeepEqual(*item2, items[1]) {
		t.Fatalf("unexpected item in list: %#v != %#v", items[1], item2)
	} else if len(items) != 2 {
		t.Fatalf("unexpected number of items in list: %d != 2", len(items))
	}
}

func TestTodoListService_List_Empty(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.TodoListService()

	items, err := s.Items()
	if err != nil {
		t.Fatal(err)
	} else if len(items) != 0 {
		t.Fatalf("unexpected number of items in list: %d != 0", len(items))
	}
}