package peer

import (
	model "github.com/t02smith/part-iii-project/toolkit/model/util"
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
		for request := range p.library.DownloadManager.RequestDownload {
			util.Logger.Debugf("Processing request for block %x", request.BlockHash)

			ps := p.findPeersWhoHaveGame(request.GameHash)
			if len(ps) == 0 {
				util.Logger.Debugf("Deferring block %x", request.BlockHash)
				p.library.DownloadManager.DeferredRequests <- request
				continue
			}

			// * at least one peer has it
			chosen := ps[0]
			chosen.SendString(generateBLOCK(request.GameHash, request.BlockHash))

			chosenPD := p.GetPeer(chosen)

			chosenPD.Lock()
			chosenPD.sentRequests[request] = model.Void{}
			chosenPD.Unlock()
		}
		util.Logger.Info("stopped listening to incoming download requests")
	}()
}

/*
queues all deferred requests
*/
func LoadDeferredRequests() {
	manager := Peer().Library().DownloadManager
	cached := manager.DeferredRequests

	go func() {
		util.Logger.Infof("loading %d deferred requests", len(cached))
		size := len(cached)

		for i := 0; i < size; i++ {
			manager.RequestDownload <- <-cached
		}
		util.Logger.Infof("deferred requests loaded")
	}()

}
