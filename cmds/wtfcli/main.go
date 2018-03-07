package main

import (
	"github.com/kristoiv/wtf"
	"github.com/kristoiv/wtf/grpc"
)

func main() {
	client := grpc.NewClient()
	cui := &CUI{
		TodoListService: client.TodoListService(),
		items:           []wtf.Item{wtf.Item{Title: "Test 1"}},
	}
	cui.loop()
}
