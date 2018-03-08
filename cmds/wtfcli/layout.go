package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func (cui *CUI) layoutManager(g *gocui.Gui) error {
	cui.updateItems()

	dy := 1
	maxX, _ := g.Size()
	count := len(cui.items)
	width := cui.percentageToWidth(0.9, maxX, 80)
	if count != 0 {
		if v, err := g.SetView("list", maxX/2-int(width/2), dy, maxX/2+int(width/2), 2+count); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
		} else {
			v.Clear()
			if len(cui.items) == 0 {
				v.Frame = false
			} else {
				v.Frame = true
			}

			v.Title = "TODOs"
			// v.Highlight = true
			//v.Wrap = true
			v.SelBgColor = gocui.ColorGreen
			v.SelFgColor = gocui.ColorBlack
			for _, item := range cui.items {
				fmt.Fprintf(v, "%s (^D=Delete)\n", item.Title)
			}
		}
		dy = 2 + count + 2
	} else {
		g.DeleteView("list")
	}

	if v, err := g.SetView("compose", maxX/2-int(width/2), dy, maxX/2+int(width/2), 2+count+4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Compose new item"
		v.Editable = true
		v.Wrap = true
		fmt.Fprint(v, " ")
		v.SetCursor(1, 0)
	}

	return nil
}
