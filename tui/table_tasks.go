package tui

import (
	"sort"

	"github.com/gdamore/tcell/v2"
	"github.com/luthermonson/go-proxmox"
	"github.com/rivo/tview"
)

func (table n00bTable[T, U]) Updatee(data T) {

	switch tableData := any(data); tableData.(type) {
	case []proxmox.Task:
		table.setTaskTableData(tableData.([]proxmox.Task))
	case []proxmox.Cluster:
	}
}

func (table n00bTable[_, _]) setTaskTableData(data []proxmox.Task) {
	init := false
	if table.GetRowCount() == 1 {
		init = true
	}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].StartTime.After(data[j].StartTime)
	})
	for i, t := range data {
		table.app.QueueUpdateDraw(func() {
			color := tcell.ColorDefault
			if t.Status != "OK" {
				color = tcell.ColorRed
			}

			table.SetCell(i+1, 1, tview.NewTableCell(t.StartTime.Format("02 Jan 06 15:04 MST")).SetTextColor(color))
			table.SetCell(i+1, 2, tview.NewTableCell(t.EndTime.Format("02 Jan 06 15:04 MST")).SetTextColor(color))
			table.SetCell(i+1, 3, tview.NewTableCell(t.Node).SetTextColor(color))
			table.SetCell(i+1, 4, tview.NewTableCell(t.User).SetTextColor(color))
			table.SetCell(i+1, 5, tview.NewTableCell(t.Type).SetTextColor(color))
			table.SetCell(i+1, 6, tview.NewTableCell(t.Status).SetTextColor(color))
		})
	}
	offset, _ := table.GetOffset()
	if init || (!init && offset < 2) {
		table.app.QueueUpdateDraw(func() {
			table.ScrollToBeginning()
		})
	}

}
func (app *n00bctl) GetTaskTable() n00bTable[[]proxmox.Task, any] {
	t := tview.NewTable().SetSeparator('|').SetFixed(1, 1)
	table := n00bTable[[]proxmox.Task, any]{t, app, TasksTableColumns, nil, nil, nil}
	table.rowFunc = table.setTaskTableData
	for index, column := range TasksTableColumns {
		table.SetCellSimple(0, index, column.header)
	}
	update := func() {
		msgCh := app.TasksBroker.Subscribe()
		for resss := range msgCh {
			table.Updatee(resss)
		}
	}

	go update()
	return table
}

var TasksTableColumns = n00bTableHeader{
	1: {"Start Time", "startTime"},
	2: {"End Time", "endTime"},
	3: {"Node", "node"},
	4: {"User Name", "user"},
	5: {"Description", "desc"},
	6: {"Status", "status"},
}
