package controller

import (
	"sort"

	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/util"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// get information about peers
func (a *Controller) GetPeerInformation() []ControllerPeerData {
	var ps []ControllerPeerData

	p := peer.Peer()
	for _, data := range p.GetPeers() {
		var games []string
		for g := range data.Library {
			game := p.Library().GetOwnedGame(g)
			if game == nil {
				continue
			}

			games = append(games, game.Title)
		}

		ps = append(ps, ControllerPeerData{
			Hostname: data.Hostname,
			Port:     data.Port,
			Library:  games,
		})
	}

	sort.Slice(ps, func(i, j int) bool {
		return len(ps[i].Library) > len(ps[j].Library)
	})

	return ps
}

// form a connection with a new peer
func (c *Controller) ConnectToPeer(hostname string, port uint) {
	err := peer.Peer().ConnectToPeer(hostname, port)
	if err != nil {
		c.controllerErrorf("Error connecting to peer %s", err)
		return
	}

	util.Logger.Infof("Connected to peer %s:%d", hostname, port)
	runtime.EventsEmit(c.ctx, "new-peer")
}
