package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (app *n00bctl) initTUILayout() {
	app.Layout.Panes["subTreeFlex"] = tview.NewFlex()
	rootDir := app.Cluster.Name
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorRed).SetSelectedFunc(func() {
		app.getContentSummary().Clear()
		app.getContentSummary().AddItem(app.getClusterSummary(), 0, 5, false)
	})
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)
	app.Layout.Panes["mainTree"] = tree
	app.Layout.Panes["subTree"] = app.prepSubTree()

	app.prepareMainTree()
	tree.SetChangedFunc(func(node *tview.TreeNode) {
		if node.GetLevel() == 2 && node.GetReference() == app.getNodeRefFromTree() {
			subt := app.GetNodeSubTree(node.GetText())
			app.Layout.Panes["subTree"] = subt
			app.subTreeFlex().Clear().AddItem(subt, 0, 3, true)
		} else {
			subt := app.GetNodeSubTree(node.GetText())
			app.Layout.Panes["subTree"] = subt
			app.subTreeFlex().Clear().AddItem(subt, 0, 3, true)
		}

	})
	tree.SetSelectedFunc(func(node *tview.TreeNode) { app.SetFocus(app.getSubTree()) })

	go app.UpdateAllData()

}
