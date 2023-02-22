package profile1

import (
	"fmt"

	"github.com/spf13/viper"
)

/*

Profile 1:

Features:
- listen-and-respond only peer
- will upload a game to ETH and seed it
- ideal connection

*/

func Run() {

	// * setup config
	viper.AddConfigPath("./test/profiles/1")
	viper.ReadInConfig()
	fmt.Println(viper.AllSettings())

	// * start peer

	// * upload game

	// * listen for connections

}
