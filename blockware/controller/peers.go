package controller

import (
	"fmt"

	"github.com/t02smith/part-iii-project/toolkit/model/net"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

// get information about peers
func (a *Controller) GetPeerInformation() []string {
	ps := []string{}

	p := net.Peer()
	for _, data := range p.GetPeers() {
		ps = append(ps, fmt.Sprintf("%s:%d", data.Hostname, data.Port))
	}

	return ps
}

// form a connection with a new peer
func (c *Controller) ConnectToPeer(hostname string, port uint) {
	err := net.Peer().ConnectToPeer(hostname, port)
	if err != nil {
		c.controllerErrorf("Error connecting to peer %s", err)
		return
	}

	util.Logger.Infof("Connected to peer %s:%d", hostname, port)
}
