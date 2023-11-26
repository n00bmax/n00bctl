package tui

import (
	"context"
	"io"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (app *n00bctl) getTermView() *tview.TextView {
	return app.Layout.Panes["termview"].(*tview.TextView)
}
func (app *n00bctl) getRightPane() *tview.Flex {
	return app.Layout.Panes["rightPane"].(*tview.Flex)
}
func (app *n00bctl) prepRightPane() {

	tView := tview.NewTextView().SetDynamicColors(true)
	app.Layout.Panes["termview"] = tView
	tView.Clear()
	app.Layout.Panes["rightPane"] = tview.NewFlex().
		AddItem(app.prepResizeButton(), 1, 1, false).
		AddItem(tView, 0, 1, false)
	app.prepRightPaneWriter()
}

func (app *n00bctl) prepRightPaneWriter() {
	w := tview.ANSIWriter(app.getTermView())
	app.Writers.RightPane = w
}

func (app *n00bctl) prepResizeButton() *tview.Button {
	button := tview.NewButton("<")
	button.
		SetActivatedStyle(tcell.Style{}).
		SetDisabledStyle(tcell.Style{}).
		SetStyle(tcell.StyleDefault.Background(tcell.ColorDefault)).
		SetSelectedFunc(func() {
			if button.GetLabel() == "<" {
				app.getContentPane().ResizeItem(app.getRightPane(), 0, 6)
				button.SetLabel(">")
			} else {
				app.getContentPane().ResizeItem(app.getRightPane(), 0, 3)
				button.SetLabel("<")
			}
		})
	return button
}

func (app *n00bctl) addRightPaneWriter(w io.Writer) {
	w2 := io.MultiWriter((app.Writers.RightPane), w)
	app.Writers.RightPane = w2
}

func (app *n00bctl) testTf() {
	app.SetStdout(app.Writers.RightPane)
	app.Plan(context.Background())

}
