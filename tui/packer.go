package tui

import (
	"os/exec"
)

func (app *n00bctl) testPacker() {
	cmd := exec.Command("packer", "build", "terrible/packer/")
	cmd.Stdout = app.Writers.RightPane
	cmd.Stderr = app.Writers.RightPane
	go cmd.Run()
	// if err != nil {
	// 	app.getHeader().SetTitle(err.Error())
	// }
}
