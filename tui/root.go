package tui

import (
	"github.com/luthermonson/go-proxmox"
	"github.com/rivo/tview"
)

func Init(client *proxmox.Client) {
	cluster, _ := client.Cluster()
	n00bui := initTUI(cluster)
	version, _ := client.Version()

	n00bui.Terraform = InitTerraform()
	n00bui.prepRightPane()

	n00bui.initTUILayout()

	n00bui.Panes["content"] = tview.NewFlex()
	n00bui.Panes["clusterSummary"] = n00bui.GetClusterSummaryLayout()
	n00bui.Panes["contentSummary"] = tview.NewFlex().AddItem(n00bui.getClusterSummary(), 0, 1, false)

	n00bui.getContentPane().
		AddItem(n00bui.getMainTree(), 0, 1, true).
		AddItem(n00bui.subTreeFlex(), 0, 1, true).
		AddItem(n00bui.getContentSummary(), 0, 5, true).
		AddItem(n00bui.getRightPane(), 0, 3, true)

	n00bui.Panes["header"] = n00bui.prepHeader()
	n00bui.getHeader().SetBorder(true).SetTitle("Proxmox Virtual Environment " + version.Version)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(n00bui.getHeader(), 3, 1, false).
			AddItem(n00bui.getContentPane(),
				0, 8, true).
			AddItem(n00bui.GetTaskTable(), 0, 3, true),
			0, 2, true)
	if err := n00bui.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
func (app *n00bctl) getContentSummary() *tview.Flex {
	return app.Layout.Panes["contentSummary"].(*tview.Flex)
}
func (app *n00bctl) getClusterSummary() *tview.Flex {
	return app.Layout.Panes["clusterSummary"].(*tview.Flex)
}
func initTUI(cluster *proxmox.Cluster) *n00bctl {

	n00bui := &n00bctl{
		Application: tview.NewApplication(),
		Cluster:     cluster,
		Layout:      Layout{Panes: make(map[string]any)},
		Brokers: Brokers{
			NewBroker[map[string]proxmox.ClusterResources](),
			NewBroker[[]proxmox.Task](),
			NewBroker[proxmox.NodeStatuses](),
			NewBroker[proxmox.Cluster](),
		},
	}

	return n00bui
}
