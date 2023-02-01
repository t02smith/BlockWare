package web

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/lib"
	"github.com/t02smith/part-iii-project/toolkit/lib/net"
	"github.com/t02smith/part-iii-project/toolkit/web/controller"
)

func StartFrontend() {

	err := startPeer()
	if err != nil {
		lib.Logger.Panicf("Failed to start peer %s", err)
	}

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/css", "web/static/css")

	r.LoadHTMLGlob("web/templates/**/*")
	controller.Router(r)

	lib.Logger.Info("server started")
	r.Run()
}

func startPeer() error {
	_, err := net.StartPeer(
		"localhost",
		6748,
		viper.GetString("games.installFolder"),
		viper.GetString("meta.directory"),
	)

	return err
}
