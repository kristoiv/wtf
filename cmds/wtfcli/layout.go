package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func (cui *CUI) layoutManager(g *gocui.Gui) error {
	maxX, _ := g.Size()
	width := cui.percentageToWidth(0.9, maxX, 200)
	left := int(float64(maxX/2) - float64(width/2))
	right := left + int(width)

	countItems := len(cui.items)

	dy := 1

	if countItems > 0 {
		if v, err := g.SetView("list", left, dy, right, 2+countItems); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			if err := cui.drawItemView(g, v); err != nil {
				return err
			}
		}
		dy = 2 + countItems + 2
	}

	if v, err := g.SetView("compose", left, dy, right, dy+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if err := cui.drawComposeView(g, v); err != nil {
			return err
		}

		if countItems > 0 {
			g.SetCurrentView("list")
		} else {
			g.SetCurrentView("compose")
		}
	}

	return nil
}

func (cui *CUI) drawItemView(g *gocui.Gui, v *gocui.View) error {
	v.Title = "Todo List"
	for _, item := range cui.items {
		fmt.Fprintf(v, "%s (^D=Remove, Space=ToggleChecked)\n", item.Title)
	}
	return nil
}

func (cui *CUI) drawComposeView(g *gocui.Gui, v *gocui.View) error {
	v.Title = "Compose Todo Item"
	v.Editable = true
	return nil
}

func (cui *CUI) percentageToWidth(percentage float64, maxX, maxWidth int) float64 {
	out := float64(maxX) * percentage
	if out > float64(maxWidth) {
		return float64(maxWidth)
	}
	return out
}
