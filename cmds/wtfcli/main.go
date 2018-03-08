package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/kristoiv/wtf"
	"github.com/kristoiv/wtf/grpc"
)

func main() {
	client := grpc.NewClient()
	todoListService := &wtf.TodoListServiceCache{TodoListService: client.TodoListService()}
	cui := &CUI{TodoListService: todoListService}

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.SetManagerFunc(cui.layoutManager)

	cui.updateItems()
	cui.setupKeybindings(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func (cui *CUI) updateItems() {
	if items, err := cui.TodoListService.Items(); err != nil {
		log.Panicln(err)
	} else {
		cui.items = items
	}
}

func (cui *CUI) percentageToWidth(percentage float64, maxX, maxWidth int) float64 {
	out := float64(maxX) * percentage
	if out > float64(maxWidth) {
		return float64(maxWidth)
	}
	return out
}
