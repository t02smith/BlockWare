package peer

import (
	"crypto/sha256"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
)

/*

function: loadDeferredRequeust
purpose: move deferred requests into the main request channel

? Test cases
success
	#1 => no deferrred requests
	#2 => one deferred request
	#3 => many deferred requests

*/

func TestLoadDeferredRequests(t *testing.T) {
	manager := Peer().Library().DownloadManager

	t.Run("success", func(t *testing.T) {
		t.Run("no deferred requests", func(t *testing.T) {
			loadDeferredRequests()
			time.Sleep(10 * time.Millisecond)

			assert.Zero(t, len(manager.RequestDownload))
			assert.Zero(t, len(manager.DeferredRequests))
		})

		requests := []games.DownloadRequest{
			{GameHash: sha256.Sum256([]byte("hello")), BlockHash: sha256.Sum256([]byte("there"))},
			{GameHash: sha256.Sum256([]byte("tom")), BlockHash: sha256.Sum256([]byte("smith"))},
			{GameHash: sha256.Sum256([]byte("pls")), BlockHash: sha256.Sum256([]byte("work"))},
		}

		t.Run("one deferred request", func(t *testing.T) {
			go func() {
				manager.DeferredRequests <- requests[0]
			}()

			time.Sleep(10 * time.Millisecond)
			loadDeferredRequests()
			time.Sleep(10 * time.Millisecond)

			assert.Equal(t, 1, len(manager.RequestDownload))
			assert.Zero(t, len(manager.DeferredRequests))

			req := <-manager.RequestDownload
			assert.Equal(t, requests[0].BlockHash, req.BlockHash)
			assert.Equal(t, requests[0].GameHash, req.GameHash)
		})

		t.Run("many deferred requests", func(t *testing.T) {
			go func() {
				for _, r := range requests {
					manager.DeferredRequests <- r
				}
			}()

			time.Sleep(10 * time.Millisecond)
			loadDeferredRequests()
			time.Sleep(10 * time.Millisecond)

			assert.Equal(t, 3, len(manager.RequestDownload))
			assert.Zero(t, len(manager.DeferredRequests))

			for _, r := range requests {
				req := <-manager.RequestDownload
				assert.Equal(t, r.BlockHash, req.BlockHash)
				assert.Equal(t, r.GameHash, req.GameHash)
			}
		})
	})
}
