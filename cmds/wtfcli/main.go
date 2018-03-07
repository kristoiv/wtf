package main

import (
	"github.com/kristoiv/wtf"
	"github.com/kristoiv/wtf/grpc"
)

func main() {
	client := grpc.NewClient()
	todoListService := &wtf.TodoListServiceCache{TodoListService: client.TodoListService()}
	cui := &CUI{TodoListService: todoListService}
	cui.loop()
}
