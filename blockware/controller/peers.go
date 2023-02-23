package controller

import (
	"fmt"

	"github.com/t02smith/part-iii-project/toolkit/model/net"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

func (a *Controller) GetPeerInformation() []string {
	ps := []string{}

	p := net.Peer()
	for _, data := range p.GetPeers() {
		ps = append(ps, fmt.Sprintf("%s:%d", data.Hostname, data.Port))
	}

	return ps
}

func (a *Controller) ConnectToPeer(hostname string, port uint) {
	err := net.Peer().ConnectToPeer(hostname, port)
	if err != nil {
		util.Logger.Errorf("Error connecting to peer %s", err)
	}
	util.Logger.Infof("Connected to peer %s:%d", hostname, port)
}
