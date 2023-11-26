package tui

import (
	"fmt"

	"github.com/hako/durafmt"
	"github.com/luthermonson/go-proxmox"
	"github.com/rivo/tview"
)

func (app *n00bctl) GetClusterNodesSummaryTable() n00bTable[[]proxmox.Cluster, any] {
	t := tview.NewTable().SetSeparator('|').SetFixed(1, 1)
	table := n00bTable[[]proxmox.Cluster, any]{t, app, NodesSummaryTableColumns, nil, nil, nil}
	for index, column := range NodesSummaryTableColumns {
		table.SetCellSimple(0, index, column.header)
	}
	update := func() {
		msgCh := app.ClusterStatusBroker.Subscribe()
		// msgCh2 := app.NodeStatusBroker.Subscribe()

		for {
			select {
			case data := <-msgCh:
				app.QueueUpdateDraw(func() {
					for i, t := range data.Nodes {
						table.SetCell(i+1, 2, tview.NewTableCell(t.IP))
						table.SetCell(i+1, 5, tview.NewTableCell(t.Status))
						data.Resources()
						duration, _ := durafmt.ParseString(fmt.Sprintf("%vs", t.Uptime))
						table.SetCell(i+1, 3, tview.NewTableCell(duration.String()))
						table.SetCell(i+1, 1, tview.NewTableCell(t.Name))
						table.SetCell(i+1, 4, tview.NewTableCell(t.Type))

					}
				})
				// case data := <-msgCh2:
				// 	app.QueueUpdateDraw(func() {
				// 		for i, t := range data {
				// 			duration, _ := durafmt.ParseString(fmt.Sprintf("%vs", t.Uptime))
				// 			table.SetCell(i+1, 3, tview.NewTableCell(duration.String()))
				// 			table.SetCell(i+1, 1, tview.NewTableCell(t.Node))
				// 			table.SetCell(i+1, 4, tview.NewTableCell(t.Type))
				// 		}
				// 	})
			}
		}
	}
	go update()

	return table
}

var NodesSummaryTableColumns = n00bTableHeader{
	1: {"Node", "node"},
	2: {"IP", "ip"},
	3: {"Uptime", "uptime"},
	4: {"ID", "id"},
	5: {"Status", "status"},
}
