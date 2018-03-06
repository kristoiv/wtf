package main

import (
	"log"

	"github.com/kristoiv/wtf/grpc"
)

func main() {
	client := grpc.NewClient()
	items, err := client.TodoListService().Items()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", items)
}
