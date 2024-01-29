package tui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"k8s.io/klog/v2"
)

func (app *n00bctl) GetNodeSubTree(node string) *tview.TreeView {
	root := tview.NewTreeNode(node).
		SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root).SetGraphics(false)
	sNode := tview.NewTreeNode("Summary").SetColor(tcell.ColorBlue)
	sNode.SetSelectedFunc(func() {
		// app.Panes["contentSummary"] = app.GetNodeSummaryLayout()
		app.getContentSummary().Clear()
		app.getContentSummary().AddItem(app.GetNodeSummaryLayout(), 0, 5, false)
	})
	root.AddChild(sNode)
	pNode := tview.NewTreeNode("PCIe Devices").SetColor(tcell.ColorBlue)
	root.AddChild(pNode)
	tree.SetDoneFunc(func(key tcell.Key) { app.SetFocus(app.getMainTree()) })
	return tree
}
func (app *n00bctl) getMainTree() *tview.TreeView {
	return app.Layout.Panes["mainTree"].(*tview.TreeView)
}
func (app *n00bctl) prepareMainTree() {
	root := app.getMainTree().GetRoot()
	nodes, _ := app.Client.Nodes()
	tNode := tview.NewTreeNode("Nodes").SetColor(tcell.ColorBlue)
	root.AddChild(tNode)
	tVMs := tview.NewTreeNode("Virtual Machine").
		SetColor(tcell.ColorBlue)
	root.AddChild(tVMs)
	tStorage := tview.NewTreeNode("Storage").
		SetColor(tcell.ColorOrange)
	root.AddChild(tStorage)
	tNets := tview.NewTreeNode("SDN").
		SetColor(tcell.ColorOrange)
	root.AddChild(tNets)
	cRes, _ := app.Cluster.Resources()
	for _, v := range cRes {

		switch v.Type {
		case "qemu":
		case "storage":
		case "sdn":
			_, v.ID, _ = strings.Cut(v.ID, "/")
			nodeName, netName, _ := strings.Cut(v.ID, "/")

			// _, nodeName, _ = strings.Cut(nodeName, "/")
			sdnText := fmt.Sprintf("%v (%v)", netName, nodeName)

			tNets.AddChild(tview.NewTreeNode(sdnText))

		}
	}

	for _, node := range nodes {
		tNode.AddChild(tview.NewTreeNode(fmt.Sprintf("%+v", node.Node)).
			SetColor(tcell.ColorOrange).SetReference(tNode))
		n, err := app.Client.Node(node.Node)
		if err != nil {
			continue
		}
		storages, err := n.Storages()
		if err != nil {
			klog.Exit(err)
		}
		sort.SliceStable(storages, func(i, j int) bool {
			return storages[i].Name < storages[j].Name
		})
		for _, storage := range storages {
			storText := fmt.Sprintf("%v (%v)", storage.Name, storage.Node)

			tStorage.AddChild(tview.NewTreeNode(storText))
		}
		vms, _ := n.VirtualMachines()
		sort.SliceStable(vms, func(i, j int) bool {
			return vms[i].VMID < vms[j].VMID
		})
		for _, vm := range vms {
			vmText := fmt.Sprintf("%v (%v)", vm.VMID, vm.Name)
			tVMs.AddChild(tview.NewTreeNode(vmText))

		}
	}
}
