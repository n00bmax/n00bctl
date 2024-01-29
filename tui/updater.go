package tui

import (
	"time"

	"github.com/luthermonson/go-proxmox"
)

func (app *n00bctl) UpdateResources() {
	mres := make(map[string]proxmox.ClusterResources)
	ress, _ := app.Resources()
	for _, v := range ress {
		mres[v.Type] = append(mres[v.Type], v)
	}

	app.ResourcesBroker.Publish(mres)
}

func (app *n00bctl) UpdateAllData() {
	go app.ResourcesBroker.Start()
	go app.TasksBroker.Start()
	go app.NodeStatusBroker.Start()
	go app.ClusterStatusBroker.Start()

	tick := time.NewTicker(3 * time.Second)
	for ; true; <-tick.C {
		go app.UpdateResources()
		go app.UpdateTasks()
		go app.UpdateCluster()
		go app.UpdateNodeStatuses()
	}

}
func (app *n00bctl) UpdateTasks() {
	cTasks, _ := app.Tasks()
	app.TasksBroker.Publish(cTasks)
}

func (app *n00bctl) UpdateNodeStatuses() {
	cNodes, _ := app.Client.Nodes()
	app.NodeStatusBroker.Publish(cNodes)
}
func (app *n00bctl) UpdateCluster() {
	cluster, _ := app.Client.Cluster()
	res, _ := cluster.Resources()
	for i, v := range cluster.Nodes {
		for _, v2 := range res {
			if v.ID == v2.ID {
				cluster.Nodes[i].Uptime = v2.Uptime
			}
		}
	}
	app.ClusterStatusBroker.Publish(*cluster)
}
