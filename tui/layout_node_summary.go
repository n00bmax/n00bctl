package tui

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"

	"github.com/rivo/tview"
)

func (app *n00bctl) GetNodeSummaryLayout() *tview.Flex {

	gauge := tvxwidgets.NewActivityModeGauge()
	gauge.SetTitle("activity mode gauge")
	gauge.SetPgBgColor(tcell.ColorOrange)
	gauge.SetRect(10, 4, 50, 3)
	gauge.SetBorder(true)

	update := func() {
		tick := time.NewTicker(500 * time.Millisecond)
		for range tick.C {
			gauge.Pulse()
			app.Draw()
		}
	}
	go update()

	mainFlex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(gauge, 0, 1, false).AddItem(app.CPUGauge(), 0, 1, false).AddItem(app.MEMGauge(), 0, 1, false),
			0, 2, false)

	return mainFlex
}
