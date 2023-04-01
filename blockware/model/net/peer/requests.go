package peer

import (
	"time"

	"github.com/t02smith/part-iii-project/toolkit/util"
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

			resend := false
			var chosenPD *peerData = nil

			for _, peerCon := range ps {
				peerData := p.GetPeer(peerCon)

				peerData.lock.Lock()
				if sentAt, ok := peerData.sentRequests[request]; ok {
					// ? we've already sent them this request
					if time.Since(sentAt).Seconds() >= 5 {
						// ? request time out => send again
						util.Logger.Info("Request timeout. Sending to new Peer.")
						delete(peerData.sentRequests, request)

					} else {
						// ? request still to be timed out
						resend = false
						peerData.lock.Unlock()
						break
					}
				} else {
					// ? request not sent to this peer
					if chosenPD == nil {
						resend = true
						chosenPD = peerData
					}
				}
				peerData.lock.Unlock()
			}

			if !resend {
				continue
			}

			// * at least one peer has it
			chosenPD.Peer.SendString(generateBLOCK(request.GameHash, request.BlockHash))

			chosenPD.Lock()
			chosenPD.sentRequests[request] = time.Now()
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
