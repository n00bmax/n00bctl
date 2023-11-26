package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (app *n00bctl) getContentPane() *tview.Flex {
	return app.Layout.Panes["content"].(*tview.Flex)
}
func (app *n00bctl) getHeader() *tview.Flex {
	return app.Layout.Panes["header"].(*tview.Flex)
}
func (app *n00bctl) prepHeader() *tview.Flex {
	dropdown := tview.NewDropDown().
		SetLabel("Select an option (hit Enter): ").
		SetOptions([]string{"First", "Second", "Third", "Fourth", "Fifth"}, nil)
	panebut := tview.NewButton("Deploy Cluster").SetSelectedFunc(func() { go app.testTf() }).SetStyle(tcell.Style{})
	panebut2 := tview.NewButton("Create Cloud-init template").SetSelectedFunc(func() { app.testPacker() })

	return tview.NewFlex().
		AddItem(panebut2, 0, 1, false).
		AddItem(panebut, 0, 1, false).
		// AddItem(panebut2, 0, 1, false).
		AddItem(dropdown, 0, 1, false)
}
