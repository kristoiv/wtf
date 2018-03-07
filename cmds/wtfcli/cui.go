package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"code.google.com/p/go-uuid/uuid"
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

	g.Cursor = true
	g.SetManagerFunc(cui.layoutManager)

	cui.setupKeybindings(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func (cui *CUI) layoutManager(g *gocui.Gui) error {
	maxX, _ := g.Size()

	count := len(cui.items)
	width := cui.percentageToWidth(0.9, maxX, 80)
	if count != 0 {
		if v, err := g.SetView("list", maxX/2-int(width/2), 1, maxX/2+int(width/2), 2+count); err != nil {
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
	} else {
		g.DeleteView("list")
	}

	if v, err := g.SetView("compose", maxX/2-int(width/2), 2+count+2, maxX/2+int(width/2), 2+count+4); err != nil {
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

func (cui *CUI) setupKeybindings(g *gocui.Gui) {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, cui.quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, cui.nextView); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("list", gocui.KeyArrowDown, gocui.ModNone, cui.cursorDown); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("list", gocui.KeyArrowUp, gocui.ModNone, cui.cursorUp); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("list", gocui.KeyCtrlD, gocui.ModNone, cui.deleteItem); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("list", gocui.KeySpace, gocui.ModNone, cui.toggleChecked); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("compose", gocui.KeyEnter, gocui.ModNone, cui.store); err != nil {
		log.Panicln(err)
	}

	// if err := g.SetKeybinding("compose", gocui.KeyTab, gocui.ModNone, cui.nextView); err != nil {
	// 	log.Panicln(err)
	// }

	// if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, cui.getLine); err != nil {
	// 	log.Panicln(err)
	// }

	// if err := g.SetKeybinding("msg", gocui.KeyEnter, gocui.ModNone, cui.delMsg); err != nil {
	// 	log.Panicln(err)
	// }
}

func (cui *CUI) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (cui *CUI) toggleChecked(g *gocui.Gui, v *gocui.View) error {
	_, idx := v.Cursor()
	item := cui.items[idx]
	item.Checked = !item.Checked
	cui.items = append(cui.items[:idx], append([]wtf.Item{item}, cui.items[idx+1:]...)...)
	return nil
}

func (cui *CUI) deleteItem(g *gocui.Gui, v *gocui.View) error {
	if len(cui.items) == 0 {
		return nil
	}
	_, idx := v.Cursor()
	cui.items = append(cui.items[:idx], cui.items[idx+1:]...)
	if len(cui.items) == 0 {
		cui.nextView(g, v)
	}
	return nil
}

func (cui *CUI) store(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()

	line, err := v.Line(cy)
	if err != nil {
		return nil
	}

	line = strings.TrimPrefix(line, " ")
	if line == "" {
		return nil
	}

	cui.items = append(cui.items, wtf.Item{
		ID:      wtf.ItemID(uuid.New()),
		Title:   line,
		Created: time.Now().UTC(),
	})

	v.Clear()
	fmt.Fprint(v, " ")
	v.SetCursor(1, cy)

	cui.nextView(g, v)
	return nil
}

func (cui *CUI) nextView(g *gocui.Gui, v *gocui.View) error {
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

func (cui *CUI) cursorUp(g *gocui.Gui, v *gocui.View) error {
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

func (cui *CUI) cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		//fmt.Fprintf(os.Stderr, "              %d,%d", oy, cy)
		if err := v.SetCursor(cx, cy+1); err != nil && cy < len(cui.items)-1 {
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func (cui *CUI) percentageToWidth(percentage float64, maxX, maxWidth int) float64 {
	out := float64(maxX) * percentage
	if out > float64(maxWidth) {
		return float64(maxWidth)
	}
	return out
}
