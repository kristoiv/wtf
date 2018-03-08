package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/kristoiv/wtf"
)

type CUI struct {
	TodoListService wtf.TodoListService
	items           []wtf.Item
}

func (cui *CUI) loop() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	cui.setupKeybindings(g)

	g.Cursor = true
	g.SetManagerFunc(cui.layoutManager)

	cui.updateItems()

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
