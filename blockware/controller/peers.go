package controller

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/util"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// the public hostname of the peer
func (c *Controller) GetPublicHostname() string {
	return peer.Peer().Config().PublicHostname
}

func (c *Controller) SetPublicHostname(hostname string) {
	peer.Peer().SetPublicHostname(hostname)
}

func (c *Controller) GetPort() uint {
	_, port := peer.Peer().GetServerInfo()
	return port
}

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

		var reputation string
		if data.TotalRequestsSent == 0 {
			reputation = "Unkown"
		} else {
			reputation = fmt.Sprintf("%0.2f",
				float32(data.TotalRepliesReceived)/float32(data.TotalRequestsSent),
			)
		}

		ps = append(ps, ControllerPeerData{
			Hostname:   data.Hostname,
			Port:       data.Port,
			Server:     data.Server,
			Library:    games,
			Validated:  data.Validator != nil && data.Validator.Valid(),
			Reputation: reputation,
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
		c.controllerErrorf("Error connecting to peer %s:%d", hostname, port)
		return
	}

	util.Logger.Infof("Connected to peer %s:%d", hostname, port)
	runtime.EventsEmit(c.ctx, "new-peer")
}

// connect to many peers at once
func (c *Controller) ConnectToManyPeers(lines string) {
	p := peer.Peer()
	for _, peer := range strings.Split(lines, "\n") {
		peerInfo := strings.Split(peer, ":")

		if len(peerInfo) != 2 {
			continue
		}

		// ! remove carriage return if it exists
		if peerInfo[1][len(peerInfo[1])-1] == '\r' {
			peerInfo[1] = peerInfo[1][:len(peerInfo[1])-1]
		}

		port, err := strconv.ParseUint(peerInfo[1], 10, 16)
		if err != nil {
			util.Logger.Warnf("Error parsing peer details for peer %s: %s", peer, err)
			continue
		}

		go func() {
			err = p.ConnectToPeer(peerInfo[0], uint(port))
			if err != nil {
				util.Logger.Warnf("Error conneting to peer %s:%d", peerInfo[0], port)
				return
			}
			runtime.EventsEmit(c.ctx, "new-peer")
		}()
	}
}

func (c *Controller) ConnectFromFile(filepath string) {
	ps, err := peer.LoadPeersFromFile(filepath)
	if err != nil {
		c.controllerErrorf("error reading peer file at %s", filepath)
		return
	}

	for _, newPeer := range ps {
		peerInfo := strings.Split(newPeer, ":")

		if len(peerInfo) != 2 {
			continue
		}

		// ! remove carriage return if it exists
		if peerInfo[1][len(peerInfo[1])-1] == '\r' {
			peerInfo[1] = peerInfo[1][:len(peerInfo[1])-1]
		}

		port, err := strconv.ParseUint(peerInfo[1], 10, 16)
		if err != nil {
			util.Logger.Warnf("Error parsing peer details for peer %s: %s", newPeer, err)
			continue
		}

		go func() {
			err = peer.Peer().ConnectToPeer(peerInfo[0], uint(port))
			if err != nil {
				util.Logger.Warnf("Error conneting to peer %s: %s", err)
				return
			}
			runtime.EventsEmit(c.ctx, "new-peer")
		}()
	}
}

func (c *Controller) Disconnect(hostname string, port uint) {
	p := peer.Peer()
	for con, data := range p.GetPeers() {
		if data.Hostname == hostname && data.Port == port {
			p.OnConnectionClose(con)
			return
		}
	}

	runtime.EventsEmit(c.ctx, "new-peer")
}

func (c *Controller) LoadDeferredRequests() {
	peer.LoadDeferredRequests()

}

func (c *Controller) ResendValidation(hostname string, port uint) {
	p := peer.Peer()
	for _, data := range p.GetPeers() {
		if data.Hostname == hostname && data.Port == port {
			data.ValidatePeer()
			return
		}
	}
}

func (c *Controller) RequestContributions(game string) {
	gh, err := hashStringToByte32(game)
	if err != nil {
		c.controllerError("Error parsing game hash")
		return
	}

	p := peer.Peer()
	for _, data := range p.GetPeers() {
		data.RequestContributionsReceipt(gh)
	}
}
