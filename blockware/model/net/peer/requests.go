package peer

import (
	"time"

	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

const requestTimeout time.Duration = 5 * time.Second

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
				sentAt, ok := peerData.sentRequests[request]
				peerData.lock.Unlock()

				if ok {
					// ? we've already sent them this request
					if time.Since(sentAt).Seconds() >= 5 {
						// ? request time out => send again
						util.Logger.Info("Request timeout. Sending to new Peer.")
						peerData.FailedRequest(request)

					} else {
						// ? request still to be timed out
						resend = false
						break
					}
				} else {
					// ? request not sent to this peer
					if chosenPD == nil {
						resend = true
						chosenPD = peerData
					}
				}
			}

			if !resend {
				continue
			}

			// * at least one peer has it
			chosenPD.Peer.SendString(generateBLOCK(request.GameHash, request.BlockHash))
			chosenPD.PushRequest(request)
		}
		util.Logger.Info("stopped listening to incoming download requests")
	}()

	// load timed out requests that haven't been responded to
	go func() {
		util.Logger.Debugf("Checking for timed out requests")

		for {
			time.Sleep(requestTimeout)

			var requests []games.DownloadRequest
			var expiredPeerReqs []games.DownloadRequest

			for _, pd := range Peer().peers {

				pd.Lock()
				for req, timeSent := range pd.sentRequests {
					if time.Since(timeSent) > requestTimeout {
						requests = append(requests, req)
						expiredPeerReqs = append(expiredPeerReqs, req)
					}
				}
				pd.Unlock()

				for _, req := range expiredPeerReqs {
					pd.FailedRequest(req)
				}

				expiredPeerReqs = []games.DownloadRequest{}
			}

			reqChannel := Peer().library.DownloadManager.RequestDownload
			for _, req := range requests {
				reqChannel <- req
			}
			util.Logger.Debugf("Finished loading expired requests => found %d", len(requests))
		}
	}()
}

/*
queues all deferred requests
*/
func LoadDeferredRequests() {
	manager := Peer().Library().DownloadManager
	cached := manager.DeferredRequests

	// load cached requests that haven't been sent before
	go func() {
		util.Logger.Infof("loading %d deferred requests", len(cached))
		size := len(cached)

		for i := 0; i < size; i++ {
			manager.RequestDownload <- <-cached
		}
		util.Logger.Infof("deferred requests loaded")

	}()

}
