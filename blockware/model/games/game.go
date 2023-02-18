package games

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	hashIO "github.com/t02smith/part-iii-project/toolkit/model/hash"
	"github.com/t02smith/part-iii-project/toolkit/test/testutil"
	"github.com/t02smith/part-iii-project/toolkit/util"
)

type Game struct {

	// game metadata
	Title           string   `json:"title"`
	Version         string   `json:"version"`
	ReleaseDate     string   `json:"release"`
	Developer       string   `json:"dev"`
	RootHash        [32]byte `json:"rootHash"`
	PreviousVersion [32]byte `json:"previousVersion"`

	// blockchain related
	IPFSId   string         `json:"IPFSId"`
	Price    *big.Int       `json:"price"`
	Uploader common.Address `json:"uploader"`

	// the shard data
	data *hashIO.HashTree

	Download *Download
}

// Creator

func CreateGame(title, version, releaseDate, developer, rootDir string, price *big.Int, shardSize uint, progress chan int) (*Game, error) {

	if shardSize == 0 {
		util.Logger.Errorf("shard size should be > 0")
		return nil, errors.New("invalid shard size")
	}

	// check version format
	versionMatches, err := regexp.MatchString("^(\\d+\\.)*\\d+$", version)
	if err != nil {
		util.Logger.Errorf("error matching version number to regex")
		return nil, err
	}

	if !versionMatches {
		util.Logger.Errorf("invalid version number")
		return nil, errors.New("invalid version number")
	}

	// check release date
	_, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Split(releaseDate, " m=")[0])
	if err != nil {
		util.Logger.Errorf("invalid release date given got %s", releaseDate)
		return nil, errors.New("invalid release date given")
	}

	// hash data
	tree, err := hashIO.NewHashTree(rootDir, shardSize, progress)
	if err != nil {
		return nil, err
	}

	err = tree.Hash()
	if err != nil {
		return nil, err
	}

	// game root hash
	hasher := sha256.New()
	hasher.Write([]byte(title))
	hasher.Write([]byte(version))
	hasher.Write([]byte(releaseDate))
	hasher.Write([]byte(developer))
	hasher.Write(tree.RootDir.RootHash[:])

	hash := hasher.Sum([]byte{})

	var h [32]byte
	copy(h[:], hash)

	// return value
	game := &Game{
		Title:           title,
		Version:         version,
		ReleaseDate:     releaseDate,
		Developer:       developer,
		data:            tree,
		RootHash:        h,
		IPFSId:          "",
		Price:           price,
		PreviousVersion: [32]byte{},
		Uploader:        common.HexToAddress(testutil.Accounts[0][0]),
	}

	return game, nil
}

// IO

func (g *Game) ReadHashData() error {
	dir := filepath.Join(viper.GetString("meta.directory"), "hashes")
	if len(dir) == 0 {
		return errors.New(".toolkit directory not specified")
	}

	hashFileName := fmt.Sprintf("%s-%s-%s.hash.json", g.Title, g.Version, g.Developer)
	data, err := hashIO.ReadHashTreeFromFile(filepath.Join(dir, hashFileName))
	if err != nil {
		return err
	}

	g.data = data
	return nil
}

func OutputToFile(g *Game) error {
	gameFilename := filepath.Join(viper.GetString("meta.directory"), fmt.Sprintf("%s-%s-%s.json", g.Title, g.Version, g.Developer))
	util.Logger.Infof("Outputting game data to %s", gameFilename)

	// output game metadata
	e, err := json.Marshal(g)
	if err != nil {
		return err
	}

	file, err := os.Create(gameFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(string(e))
	writer.Flush()

	// output game data
	err = g.data.OutputToFile(filepath.Join(viper.GetString("meta.directory"), "hashes", fmt.Sprintf("%s-%s-%s.hash.json", g.Title, g.Version, g.Developer)))
	if err != nil {
		return err
	}

	util.Logger.Infof("Outputted game data to %s", gameFilename)
	return nil
}

// load all the game data stored in storage
func LoadGames(gameDataLocation string) ([]*Game, error) {

	// does the directory to load from exist?
	_, err := os.Stat(gameDataLocation)
	if err != nil {
		return nil, err
	}

	// Parse the games
	games := []*Game{}

	dir, err := os.Open(gameDataLocation)
	if err != nil {
		return nil, err
	}

	gameList, err := dir.Readdirnames(0)
	dir.Close()

	if err != nil {
		return nil, err
	}

	for _, game := range gameList {
		f, err := os.Stat(filepath.Join(gameDataLocation, game))

		// we only care about json files
		if err != nil || f.IsDir() || !strings.HasSuffix(game, ".json") {
			continue
		}

		gameFile, err := os.Open(filepath.Join(gameDataLocation, game))
		if err != nil {
			continue
		}

		data, err := io.ReadAll(gameFile)
		if err != nil {
			continue
		}

		var gm Game
		err = json.Unmarshal(data, &gm)
		if err == nil {
			games = append(games, &gm)
		}
	}

	return games, nil
}

// Serialisation

// Turns a game into a base64 encoded, gzip compressed byte stream
func (g *Game) Serialise() (string, error) {

	// encode
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)

	err := e.Encode(*g)
	if err != nil {
		return "", err
	}

	// compress
	compressed := bytes.Buffer{}
	compressor := gzip.NewWriter(&compressed)

	compressor.Write(b.Bytes())
	compressor.Close()

	return base64.StdEncoding.EncodeToString(compressed.Bytes()), nil
}

// Takes a serialised game and turns it into a struct
func DeserialiseGame(data string) (*Game, error) {
	g := &Game{}

	// from base64
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	b := bytes.Buffer{}
	_, err = b.Write(decodedData)
	if err != nil {
		return nil, err
	}

	// decompress
	decompressor, err := gzip.NewReader(&b)
	if err != nil {
		return nil, err
	}
	decompressor.Close()

	decompressed := bytes.Buffer{}
	io.Copy(&decompressed, decompressor)

	// decode
	d := gob.NewDecoder(&decompressed)
	err = d.Decode(g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g1 *Game) Equals(g2 *Game) bool {
	return g1.Title == g2.Title &&
		g1.Version == g2.Version &&
		g1.Developer == g2.Developer &&
		g1.ReleaseDate == g2.ReleaseDate &&
		bytes.Equal(g1.RootHash[:], g2.RootHash[:])

}

// Fetch the shard for a given hash
func (g *Game) FetchShard(hash [32]byte) (bool, []byte, error) {
	err := g.ReadHashData()
	if err != nil {
		return false, nil, err
	}

	found, data, err := g.data.GetShard(hash)
	if err != nil {
		return false, nil, err
	}

	if !found {
		return false, nil, nil
	}

	return true, data, nil
}

func (g *Game) GetData() (*hashIO.HashTree, error) {
	if g.data == nil {
		err := g.ReadHashData()
		if err != nil {
			return nil, err
		}
	}

	return g.data, nil
}

// Get a games download
func (g *Game) GetDownload() *Download {
	return g.Download
}
