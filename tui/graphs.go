package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/luthermonson/go-proxmox"
	"github.com/navidys/tvxwidgets"
)

var clusterCPUData []float64
var clusterMEMData []float64

func (app *n00bctl) GetActivityGauge() *tvxwidgets.ActivityModeGauge {

	gauge := tvxwidgets.NewActivityModeGauge()
	gauge.SetTitle("activity mode gauge")
	gauge.SetPgBgColor(tcell.ColorOrange)
	gauge.SetRect(10, 4, 50, 3)
	// gauge.SetBorder(true)

	update := func() {
		// tick := time.NewTicker(500 * time.Millisecond)
		msgCh := app.Brokers.ResourcesBroker.Subscribe()
		for range msgCh {
			app.QueueUpdateDraw(func() {
				gauge.Pulse()
				// app.Draw()
			})
		}

	}
	go update()
	return gauge
}

func (app *n00bctl) GetGraph(resType string, nodes ...string) *tvxwidgets.Plot {
	bmLineChart := tvxwidgets.NewPlot()
	bmLineChart.SetLineColor([]tcell.Color{
		tcell.ColorSteelBlue,
		tcell.ColorGreen,
	})
	bmLineChart.SetMarker(tvxwidgets.PlotMarkerDot)
	bmLineChart.SetDotMarkerRune('\u25c9')
	bmLineChart.SetDotMarkerRune('\u25c9')
	bmLineChart.SetPlotType(tvxwidgets.PlotTypeScatter)
	bmLineChart.SetMarker(tvxwidgets.PlotMarkerBraille)
	update := func() {
		msgCh := app.Brokers.ResourcesBroker.Subscribe()
		for data := range msgCh {

			graphData := app.getGraphHandler(resType)(data, nodes...)
			app.QueueUpdateDraw(func() {
				app.RWMutex.RLock()

				bmLineChart.SetData([][]float64{
					graphData,
				})
				app.RWMutex.RUnlock()
			})

		}
	}

	go update()
	return bmLineChart
}

func (app *n00bctl) getGraphHandler(resType string) func(map[string]proxmox.ClusterResources, ...string) []float64 {
	switch resType {
	case "cpu":
		return app.UpdateCPUGraphData
	case "mem":
		return app.UpdateMEMGraphData
	}
	return nil

}

func (app *n00bctl) UpdateCPUGraphData(data map[string]proxmox.ClusterResources, nodes ...string) []float64 {
	val := 0.0
	for _, v := range data["node"] {
		// if len(nodes) != 0 && slices.Contains(nodes, v.Node) {
		// 	val += v.CPU * 100
		// } else {
		// 	val += v.CPU * 100
		// }
		val += v.CPU * 100

	}
	app.RWMutex.Lock()
	clusterCPUData = append(clusterCPUData, val)
	app.RWMutex.Unlock()
	return clusterCPUData
}
func (app *n00bctl) UpdateMEMGraphData(data map[string]proxmox.ClusterResources, nodes ...string) []float64 {
	val := 0.0
	for _, v := range data["node"] {
		if v.MaxMem != 0 {
			val += (float64(v.Mem) / float64(v.MaxMem)) * 100
		}
		// if len(nodes) != 0 && slices.Contains(nodes, v.Node) {
		// 	val += (float64(v.Mem) / float64(v.MaxMem)) * 100
		// } else {
		// 	val += (float64(v.Mem) / float64(v.MaxMem)) * 100
		// }

	}
	app.RWMutex.Lock()
	clusterMEMData = append(clusterMEMData, val)
	app.RWMutex.Unlock()
	return clusterMEMData
}
