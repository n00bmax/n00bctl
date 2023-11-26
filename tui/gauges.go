package tui

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
)

func (app *n00bctl) CPUGauge(nodes ...string) *tvxwidgets.UtilModeGauge {
	gauge := tvxwidgets.NewUtilModeGauge()
	gauge.SetLabel("cpu usage:")
	gauge.SetLabelColor(tcell.ColorLightSkyBlue)
	// gauge.SetRect(10, 4, 50, 3)
	gauge.SetWarnPercentage(65)
	gauge.SetCritPercentage(80)
	// gauge.SetBorder(true)

	update := func() {
		msgCh := app.ResourcesBroker.Subscribe()

		for resss := range msgCh {
			val := 0.0
			tCPU := 0
			for _, v := range resss["node"] {
				// if len(nodes) != 0 && slices.Contains(nodes, v.Node) {
				// 	val += v.CPU * 100
				// } else {
				// 	val += v.CPU * 100
				// }
				val += v.CPU * 100
				tCPU += int(v.MaxCPU)
			}

			app.QueueUpdateDraw(func() {
				gauge.SetLabel(fmt.Sprintf("cpus: %v", tCPU))

				gauge.SetValue(val)
			})
		}
	}

	go update()
	return gauge
}
func (app *n00bctl) MEMGauge(nodes ...string) *tvxwidgets.UtilModeGauge {
	gauge := tvxwidgets.NewUtilModeGauge()
	gauge.SetLabel("mem usage:")
	gauge.SetLabelColor(tcell.ColorGreen)
	// gauge.SetRect(10, 4, 50, 3)
	gauge.SetWarnPercentage(50)
	gauge.SetCritPercentage(80)
	// gauge.SetBorder(true)

	update := func() {

		msgCh := app.ResourcesBroker.Subscribe()

		for resss := range msgCh {
			var val float64
			var maxMem float64
			for _, v := range resss["node"] {
				if v.Status == "offline" {
					continue
				}
				// if len(nodes) != 0 && slices.Contains(nodes, v.Node) {
				// 	val += (float64(v.Mem) / float64(v.MaxMem))
				// } else {
				// 	val += (float64(v.Mem) / float64(v.MaxMem))
				// }
				val += float64(v.Mem)
				maxMem += float64(v.MaxMem)

			}
			app.QueueUpdateDraw(func() {
				tot := val / maxMem
				gauge.SetValue(tot * 100)
				gauge.SetLabel(fmt.Sprintf("mem:%v/%v", humanize.IBytes(uint64(val)), humanize.IBytes(uint64(maxMem))))
			})
			// app.Draw()

		}
	}

	go update()
	return gauge
}
func (app *n00bctl) StorageGauge(nodes ...string) *tvxwidgets.UtilModeGauge {
	gauge := tvxwidgets.NewUtilModeGauge()
	gauge.SetLabel("disk usage:")
	gauge.SetLabelColor(tcell.ColorLightSkyBlue)
	// gauge.SetRect(10, 4, 50, 3)
	gauge.SetWarnPercentage(65)
	gauge.SetCritPercentage(80)
	// gauge.SetBorder(true)

	update := func() {

		msgCh := app.ResourcesBroker.Subscribe()

		for resss := range msgCh {

			var disk float64
			var diskMax float64
			// total:=/0
			for _, v := range resss["storage"] {
				if v.Status == "offline" {
					continue
				}
				// if len(nodes) != 0 && slices.Contains(nodes, v.Node) {
				// 	disk += float64(v.Disk)
				// 	diskMax += float64(v.MaxDisk)
				// } else {
				// 	disk += float64(v.Disk)
				// 	diskMax += float64(v.MaxDisk)
				// }
				disk += float64(v.Disk)
				diskMax += float64(v.MaxDisk)
			}
			app.QueueUpdateDraw(func() {

				gauge.SetLabel(fmt.Sprintf("disk: %v", humanize.IBytes(uint64(diskMax))))
				gauge.SetValue((disk / diskMax) * 100)
			})
			// app.Draw()

		}
	}

	go update()
	return gauge
}
