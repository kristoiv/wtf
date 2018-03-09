package grpc_test

import (
	"testing"
	"time"

	"github.com/kristoiv/wtf/grpc"
)

func TestTodoListService_Add(t *testing.T) {
	client := grpc.NewClient()
	if err := client.Open(); err != nil {
		t.Fatalf("unexpected error in dialing non-existent grpc server: %s", err)
	}

	if _, err := client.TodoListService().Items(); err == nil {
		t.Fatalf("unexpected success in fetching items from a non-extisting grpc-server: %s", err)
	}

	server, mock := NewServer()
	if err := server.Open(); err != nil {
		t.Fatalf("unexpected problem starting grpc-server: %s", err)
	}

	// The fist connect attempt failed, retry waits almost a second so we do too
	time.Sleep(2 * time.Second)

	if item, err := client.TodoListService().Add("Test"); err != nil {
		t.Fatalf("unexpected error adding item: %s", err)
	} else if !mock.addCalled {
		t.Fatalf("add method never called on TodoListServiceHandler")
	} else if mock.addTitle != "Test" {
		t.Fatalf("received item title does not match expected: %q != \"Test\"", item.Title)
	} else if item.Title != "Test" {
		t.Fatalf("returned item title does not match expected: %q != \"Test\"", item.Title)
	}
}
