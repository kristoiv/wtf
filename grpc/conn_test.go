package grpc_test

import (
	"github.com/kristoiv/wtf"
	"github.com/kristoiv/wtf/grpc"
)

var _ wtf.TodoListService = &todoListMock{}

type todoListMock struct {
	addCalled        bool
	setCheckedCalled bool
	removeCalled     bool
	itemsCalled      bool

	addTitle          string
	setCheckedID      wtf.ItemID
	setCheckedChecked bool
	removeID          wtf.ItemID
}

func (s *todoListMock) Add(title string) (*wtf.Item, error) {
	s.addCalled = true
	s.addTitle = title
	return &wtf.Item{Title: title}, nil
}

func (s *todoListMock) SetChecked(id wtf.ItemID, checked bool) error {
	s.setCheckedCalled = true
	s.setCheckedID = id
	s.setCheckedChecked = checked
	return nil
}

func (s *todoListMock) Remove(id wtf.ItemID) error {
	s.removeCalled = true
	s.removeID = id
	return nil
}

func (s *todoListMock) Items() ([]wtf.Item, error) {
	s.itemsCalled = true
	return []wtf.Item{}, nil
}

func NewServer() (*grpc.Server, *todoListMock) {
	server := grpc.NewServer()
	todoListMock := &todoListMock{}
	server.TodoListServiceHandler = &grpc.TodoListServiceHandler{TodoListService: todoListMock}
	return server, todoListMock
}
