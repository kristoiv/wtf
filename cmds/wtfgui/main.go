package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jroimartin/gocui"
)

var todoItems = []string{
	" [+]\tItem A",
	" [+]\tItem B: Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet.",
	" [+]\tItem C",
	" [X]\tItem D",
	" [+]\tItem E",
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("list", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("list", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("list", gocui.KeyCtrlD, gocui.ModNone, deleteItem); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("list", gocui.KeySpace, gocui.ModNone, toggleChecked); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("compose", gocui.KeyEnter, gocui.ModNone, store); err != nil {
		log.Panicln(err)
	}

	// if err := g.SetKeybinding("compose", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
	// 	log.Panicln(err)
	// }

	// if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, getLine); err != nil {
	// 	log.Panicln(err)
	// }

	// if err := g.SetKeybinding("msg", gocui.KeyEnter, gocui.ModNone, delMsg); err != nil {
	// 	log.Panicln(err)
	// }

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, _ := g.Size()

	count := len(todoItems)
	width := percentageToWidth(0.9, maxX, 80)
	if count != 0 {
		if v, err := g.SetView("list", maxX/2-int(width/2), 1, maxX/2+int(width/2), 2+count); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
		} else {
			v.Clear()
			if len(todoItems) == 0 {
				v.Frame = false
			} else {
				v.Frame = true
			}

			v.Title = "TODOs"
			// v.Highlight = true
			//v.Wrap = true
			v.SelBgColor = gocui.ColorGreen
			v.SelFgColor = gocui.ColorBlack
			for _, item := range todoItems {
				fmt.Fprintf(v, "%s (^D=Delete)\n", item)
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

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func toggleChecked(g *gocui.Gui, v *gocui.View) error {
	_, idx := v.Cursor()
	line := todoItems[idx]
	if line[2] == '+' {
		line = " [X]" + strings.TrimPrefix(line, " [+]")
	} else {
		line = " [+]" + strings.TrimPrefix(line, " [X]")
	}
	todoItems = append(todoItems[:idx], append([]string{line}, todoItems[idx+1:]...)...)
	return nil
}

func deleteItem(g *gocui.Gui, v *gocui.View) error {
	if len(todoItems) == 0 {
		return nil
	}
	_, idx := v.Cursor()
	todoItems = append(todoItems[:idx], todoItems[idx+1:]...)
	if len(todoItems) == 0 {
		nextView(g, v)
	}
	return nil
}

func store(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()

	line, err := v.Line(cy)
	if err != nil {
		return nil
	}

	line = strings.TrimPrefix(line, " ")
	if line == "" {
		return nil
	}

	todoItems = append(todoItems, " [+] "+line)

	v.Clear()
	fmt.Fprint(v, " ")
	v.SetCursor(1, cy)

	nextView(g, v)
	return nil
}

func nextView(g *gocui.Gui, v *gocui.View) error {
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

func cursorUp(g *gocui.Gui, v *gocui.View) error {
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

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		//fmt.Fprintf(os.Stderr, "              %d,%d", oy, cy)
		if err := v.SetCursor(cx, cy+1); err != nil && cy < len(todoItems)-1 {
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func percentageToWidth(percentage float64, maxX, maxWidth int) float64 {
	out := float64(maxX) * percentage
	if out > float64(maxWidth) {
		return float64(maxWidth)
	}
	return out
}
