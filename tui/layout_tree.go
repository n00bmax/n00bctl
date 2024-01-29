package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (app *n00bctl) subTreeFlex() *tview.Flex {
	return app.Layout.Panes["subTreeFlex"].(*tview.Flex)
}

func (app *n00bctl) prepSubTree() *tview.TreeView {
	root := tview.NewTreeNode("s33rr").
		SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	tree.SetRoot(root)
	app.Layout.Panes["subTree"] = tree
	return tree
}
func (app *n00bctl) getSubTree() *tview.TreeView {
	return app.Layout.Panes["subTree"].(*tview.TreeView)
}
func (app *n00bctl) getNodeRefFromTree() *tview.TreeNode {
	for _, tn := range app.getMainTree().GetRoot().GetChildren() {
		if tn.GetLevel() == 1 && tn.GetText() == "Nodes" {
			return tn
		}
	}
	return nil
}
