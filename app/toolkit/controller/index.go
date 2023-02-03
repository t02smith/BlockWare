package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t02smith/part-iii-project/toolkit/model/net"
)

// data

type NavbarData struct {
	Text string
	Link string
}

type UserInfo struct {
	Server         string
	PeerCount      int
	GamesInLibrary int
}

//

func Router(r *gin.Engine) {
	r.GET("/", index)
}

func index(c *gin.Context) {

	// user info
	p := net.GetPeerInstance()
	host, port := p.GetServerInfo()

	c.HTML(
		http.StatusOK,
		"views/index.html",
		gin.H{
			"title": "Homepage",
			"navOptions": []NavbarData{
				{"📝 Dashboard", "/"},
				{"📖 Library", "/library"},
				{"🫂 Peers", "/peers"},
			},
			"user": UserInfo{
				Server:         fmt.Sprintf("%s:%d", host, port),
				PeerCount:      len(p.GetPeers()),
				GamesInLibrary: len(p.GetLibrary().GetGames()),
			},
			"games": p.GetLibrary().GetGames(),
		},
	)
}
