package main

import (
	"github.com/kristoiv/wtf"
)

type CUI struct {
	TodoListService wtf.TodoListService
	items           []wtf.Item
}
