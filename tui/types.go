package tui

import (
	"io"
	"sync"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/luthermonson/go-proxmox"
	"github.com/rivo/tview"
)

type n00bctl struct {
	*tview.Application
	*proxmox.Cluster
	*tfexec.Terraform
	Brokers
	Layout
	Writers
}

type Layout struct {
	paneLock *sync.RWMutex
	Panes    map[string]any
}

type Writers struct {
	RightPane io.Writer
}

type Brokers struct {
	ResourcesBroker     *Broker[map[string]proxmox.ClusterResources]
	TasksBroker         *Broker[[]proxmox.Task]
	NodeStatusBroker    *Broker[proxmox.NodeStatuses]
	ClusterStatusBroker *Broker[proxmox.Cluster]
}

type n00bTableHeader map[int]struct {
	header string
	id     string
}

type n00bTable[T, U any] struct {
	*tview.Table
	app *n00bctl
	n00bTableHeader
	data    T
	keys    U
	rowFunc func(T)
}
