package peer

import (
	"time"

	"github.com/t02smith/part-iii-project/toolkit/model"
	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

 */

const (
	maxRequestsPerPeer uint = 10
)

// listen for incoming download requests
func (p *peer) listenToDownloadRequests() {
	util.Logger.Info("Listening for incoming download requests")
	go func() {
		for len(p.peers) == 0 {
		}

		for request := range p.library.DownloadManager.RequestDownload {
			if request.Attempts > uint8(MaxAttemptsDownloadRequest) {
				// ! limit number of attempts we can make for a given download
				util.Logger.Warnf("Request for block %x cancelled after %d attempts", request.BlockHash, MaxAttemptsDownloadRequest)
				continue
			}

			util.Logger.Infof("Processing request for block %x", request.BlockHash)
			request.Attempts++

			ps := p.findPeersWhoHaveGame(request.GameHash)
			if len(ps) == 0 {
				// ! no peers have games
				p.refreshLibraries()

				p.library.DownloadManager.DeferredRequests <- request
				continue
			}

			// * at least one peer has it
			chosen := ps[0]
			chosen.SendString(generateBLOCK(request.GameHash, request.BlockHash))

			chosenPD := p.peers[chosen]
			chosenPD.SentRequests[request] = model.Void{}
			time.Sleep(100 * time.Millisecond)
		}
		util.Logger.Info("stopped listening to incoming download requests")
	}()
}

/*
queues all deferred requests
*/
func loadDeferredRequests() {
	manager := Peer().Library().DownloadManager
	cached := manager.DeferredRequests

	go func() {
		util.Logger.Infof("loading %d deferred requests", len(cached))
		defer close(cached)
		manager.DeferredRequests = make(chan games.DownloadRequest)
		for request := range cached {
			manager.RequestDownload <- request
		}
		util.Logger.Infof("deferred requests loaded")
	}()

}
