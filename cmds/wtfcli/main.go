package main

import (
	"log"

	"github.com/kristoiv/wtf/grpc"
)

func main() {
	srv := grpc.NewClient()
	items, err := srv.TodoListService().Items()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", items)
}
