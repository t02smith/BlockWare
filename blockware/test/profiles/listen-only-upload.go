package profiles

import (
	"math/big"
	"time"

	"github.com/t02smith/part-iii-project/toolkit/model/manager/games"
	"github.com/t02smith/part-iii-project/toolkit/model/net/peer"
	"github.com/t02smith/part-iii-project/toolkit/model/persistence/ethereum/library"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

/*

Profile 1:

Features:
- listen-and-respond only peer
- will upload a game to ETH and seed it
- ideal connection
*/

var listenOnlyUploadPrivateKey string = testutil.Accounts[1][1]

func listenOnlyUploadRun() {
	// if _, err := os.Stat("../../data/tmp/games/profile-2"); err != nil {
	// 	if err = testutil.GenerateLargeFolder("profile-2", "../../data/tmp/games/", 80_000_000, 500); err != nil {
	// 		util.Logger.Fatal(err)
	// 	}
	// }

	p := peer.Peer()

	// * create & upload game
	// g, err := games.CreateGame(games.NewGame{
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
	// _ = g

	// if err != nil {
	// 	util.Logger.Fatalf("Error creating game: %s", err)
	// }

	g1, err := games.CreateGame(games.NewGame{
		Title:       "website",
		Version:     "4.7.1",
		ReleaseDate: time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
		Developer:   "tcs1g20",
		RootDir:     "./games/t02smith.github.io",
		Price:       big.NewInt(150),
		ShardSize:   4194304,
		AssetsDir:   "../data/assets"},
		nil,
	)
	_ = g1

	if err != nil {
		util.Logger.Errorf("Error creating game: %s", err)
	}

	g2, err := games.CreateGame(games.NewGame{
		Title:       "Snake",
		Version:     "1.17.4",
		ReleaseDate: time.Date(2002, time.January, 10, 0, 0, 0, 0, time.UTC).String(),
		Developer:   "Tom Smith",
		RootDir:     "../../model",
		Price:       big.NewInt(10),
		ShardSize:   4194304,
		AssetsDir:   "../data/assets"},
		nil,
	)
	_ = g2

	if err != nil {
		util.Logger.Errorf("Error creating game: %s", err)
	}

	for _, _g := range []*games.Game{g1, g2} {
		p.Library().AddOrUpdateOwnedGame(_g)
		if err := library.UploadToEthereum(_g); err != nil {
			util.Logger.Fatalf("Error uploading game to ETH %s", err)
		}

		if err := games.OutputAllGameDataToFile(_g); err != nil {
			util.Logger.Warnf("Error saving game to file %s", err)
		}
	}

	// * listen for connections
	util.Logger.Info("PROFILE 1: listening")
	<-make(chan bool) // wait forever :/
}
