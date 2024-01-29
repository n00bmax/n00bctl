package tui

import (
	"github.com/rivo/tview"
)

func (app *n00bctl) GetClusterSummaryLayout() *tview.Flex {

	mainFlex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(app.GetActivityGauge(), 2, 2, false).
			AddItem(app.CPUGauge(), 1, 1, false).
			AddItem(app.MEMGauge(), 1, 1, false).
			AddItem(app.StorageGauge(), 1, 1, false).
			AddItem(app.GetGraph("mem"), 0, 1, false).
			AddItem(app.GetGraph("cpu"), 0, 1, false).
			AddItem(app.GetClusterNodesSummaryTable(), 0, 1, false),

			0, 10, false)

	return mainFlex
}
