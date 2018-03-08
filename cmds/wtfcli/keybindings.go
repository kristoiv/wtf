package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

func (cui *CUI) setupKeybindings(g *gocui.Gui) {
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, cui.nextViewAction); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, cui.quitAction); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("list", gocui.KeyArrowUp, gocui.ModNone, cui.cursorUpAction); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("list", gocui.KeyArrowDown, gocui.ModNone, cui.cursorDownAction); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("list", gocui.KeySpace, gocui.ModNone, cui.itemToggleCheckedAction); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("list", gocui.KeyCtrlD, gocui.ModNone, cui.deleteItemAction); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("compose", gocui.KeyEnter, gocui.ModNone, cui.createItemAction); err != nil {
		log.Panicln(err)
	}
}

func (cui *CUI) nextViewAction(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "compose" {
		if _, err := g.View("list"); err == nil {
			_, err := g.SetCurrentView("list")
			return err
		}
		return nil
	}
	_, err := g.SetCurrentView("compose")
	return err
}

func (cui *CUI) quitAction(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (cui *CUI) cursorUpAction(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func (cui *CUI) cursorDownAction(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil && cy < len(cui.items)-1 {
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func (cui *CUI) itemToggleCheckedAction(g *gocui.Gui, v *gocui.View) error {
	_, idx := v.Cursor()
	item := cui.items[idx]
	cui.TodoListService.SetChecked(item.ID, !item.Checked)
	return nil
}

func (cui *CUI) deleteItemAction(g *gocui.Gui, v *gocui.View) error {
	if len(cui.items) == 0 {
		return nil
	}
	_, idx := v.Cursor()
	cui.TodoListService.Remove(cui.items[idx].ID)
	if len(cui.items) == 1 {
		cui.nextViewAction(g, v)
	}
	return nil
}

func (cui *CUI) createItemAction(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	line, err := v.Line(cy)
	if err != nil {
		return nil
	}

	if line == "" {
		return nil
	}

	cui.TodoListService.Add(line)

	v.Clear()
	v.SetCursor(0, cy)
	cui.nextViewAction(g, v)
	return nil
}
