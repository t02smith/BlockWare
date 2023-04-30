package profiles

import (
	"math/big"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

Profile 1:

Features:
- listen-and-respond only peer
- will upload a game to ETH and seed it
- ideal connection

Data:

- small size project (latex-template from github)
	a small scale project that will allow for the peer to download
	from many different files.

*/

var listenOnlyPrivateKey string = testutil.Accounts[2][1]

func listenOnlyRun() {
	p := peer.Peer()

	// * create & upload game
	g1, err := games.CreateGame(games.NewGame{
		Title:       "t02smith.github.io",
		Version:     "4.7.1",
		ReleaseDate: time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
		Developer:   "tcs1g20",
		RootDir:     "../games/t02smith.github.io",
		Price:       big.NewInt(150),
		ShardSize:   4194304,
		AssetsDir:   "../data/assets"},
		nil,
	)

	if err != nil {
		util.Logger.Fatalf("Error creating game: %s", err)
	}

	// g2, err := games.CreateGame(games.NewGame{
	// 	Title:       "benchmark",
	// 	Version:     "4.7.1",
	// 	ReleaseDate: time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
	// 	Developer:   "tcs1g20",
	// 	RootDir:     "../../data/tmp/games/profile-2",
	// 	Price:       big.NewInt(150),
	// 	ShardSize:   4194304,
	// 	AssetsDir:   "../../data/assets"},
	// 	nil,
	// )

	p.Library().AddOrUpdateOwnedGame(g1)
	err = games.OutputAllGameDataToFile(g1)
	if err != nil {
		util.Logger.Warnf("Error saving game to file %s", err)
	}

	// * listen for connections
	util.Logger.Info("PROFILE 2: listening")
	<-make(chan bool) // wait forever :/
}
