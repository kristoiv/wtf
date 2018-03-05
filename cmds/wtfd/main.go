package main

import (
	"log"

	"github.com/kristoiv/wtf/bolt"
	"github.com/kristoiv/wtf/grpc"
)

func main() {
	client := bolt.NewClient()
	if err := client.Open(); err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	srv.TodoListServiceHandler = &grpc.TodoListServiceHandler{TodoListService: client.TodoListService()}
	if err := srv.Open(); err != nil {
		log.Fatalln(err)
	}

	select {}
}
