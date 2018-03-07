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

	items, _ := s.Items()
	if len(items) != 0 {
		t.Fatalf("Remove unsuccessful")
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
	} else if len(items) != 2 {
		t.Fatalf("unexpected number of items in list: %d != 2", len(items))
	}

	otherItem1 := items[0]
	otherItem2 := items[1]
	if item1.ID != otherItem1.ID {
		otherItem1, otherItem2 = otherItem2, otherItem1
	}

	if !reflect.DeepEqual(*item1, otherItem1) {
		t.Fatalf("unexpected item in list: %#v != %#v", otherItem1, item1)
	} else if !reflect.DeepEqual(*item2, otherItem2) {
		t.Fatalf("unexpected item in list: %#v != %#v", otherItem2, item2)
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
