package main

import (
	"github.com/jroimartin/gocui"
)

func (cui *CUI) layoutManager(g *gocui.Gui) error {
	maxX, _ := g.Size()
	width := cui.percentageToWidth(0.9, maxX, 200)
	left := int(float64(maxX/2) - float64(width/2))
	right := left + int(width)

	countItems := len(cui.items)

	dy := 1

	if len(cui.items) > 0 {
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
	}

	return nil
}

func (cui *CUI) drawItemView(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func (cui *CUI) drawComposeView(g *gocui.Gui, v *gocui.View) error {
	return nil
}
