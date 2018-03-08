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

	cui.updateItems()
	countItems := len(cui.items)

	dy := 1

	if countItems > 0 {
		listView, err := g.SetView("list", left, dy, right, 2+countItems)
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}
		listView.Clear()
		if err := cui.drawItemsView(g, listView); err != nil {
			return err
		}
		dy = 2 + countItems

		listViewFooter, err := g.SetView("list_footer", left, dy, right, dy+2)
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}
		listViewFooter.Clear()
		if err := cui.drawItemsFooterView(g, listViewFooter); err != nil {
			return err
		}
		dy += 4
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

func (cui *CUI) drawItemsView(g *gocui.Gui, v *gocui.View) error {
	v.Title = "Todo List"
	for _, item := range cui.items {
		checked := "+"
		if item.Checked {
			checked = "X"
		}
		fmt.Fprintf(v, " [%s] %s\n", checked, item.Title)
	}
	return nil
}

func (cui *CUI) drawItemsFooterView(g *gocui.Gui, v *gocui.View) error {
	v.Frame = false
	fmt.Fprintf(v, "(Space=Check/Uncheck item, ^D=Delete)")
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
