package games

import (
	"bufio"
	"crypto/sha256"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/viper"
	hashIO "github.com/t02smith/part-iii-project/toolkit/lib/io"
)

type Game struct {

	// game metadata
	Title       string `json:"title"`
	Version     string `json:"version"`
	ReleaseDate string `json:"release"`
	Developer   string `json:"dev"`
	RootHash    []byte `json:"rootHash"`

	// the shard data
	data *hashIO.HashTree
}

// Creator

func CreateGame(title, version, releaseDate, developer, rootDir string, shardSize uint) (*Game, error) {

	// check version format
	versionMatches, err := regexp.MatchString("^(\\d+\\.)*\\d+$", version)
	if err != nil {
		log.Println("error matching version number to regex")
		return nil, err
	}

	if !versionMatches {
		log.Println("invalid version number")
		return nil, errors.New("invalid version number")
	}

	// check release date
	_, err = time.Parse("2006-01-02 15:04:05 -0700 MST", releaseDate)
	if err != nil {
		log.Println("invalid release date given")
		return nil, errors.New("invalid release date given")
	}

	// check domain has SSL certificate
	domainCorrect, err := verifyDomain(developer)
	if err != nil {
		return nil, err
	}

	if !domainCorrect {
		return nil, errors.New("invalid domain specified")
	}

	// hash data
	tree, err := hashIO.NewHashTree(rootDir, shardSize)
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

	// return value
	game := &Game{
		Title:       title,
		Version:     version,
		ReleaseDate: releaseDate,
		Developer:   developer,
		data:        tree,
		RootHash:    hash,
	}

	return game, nil
}

func verifyDomain(domain string) (bool, error) {

	conf := &tls.Config{}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:443", domain), conf)
	if err != nil {
		return false, err
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		return false, nil
	}

	return true, nil
}

// IO

func OutputToFile(g *Game) error {
	gameFilename := filepath.Join(viper.GetString("meta.directory"), fmt.Sprintf("%s-%s-%s.json", g.Title, g.Version, g.Developer))
	fmt.Printf("Outputting game data to %s\n", gameFilename)

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
	err = g.data.OutputToFile(filepath.Join(viper.GetString("meta.hashes.directory"), fmt.Sprintf("%s-%s-%s.hash.json", g.Title, g.Version, g.Developer)))
	if err != nil {
		return err
	}

	return nil
}

// load all the game data stored in storage
func LoadGames() ([]*Game, error) {
	dirName := viper.GetString("meta.directory")

	// does the directory to load from exist?
	_, err := os.Stat(dirName)
	if err != nil {
		return nil, err
	}

	// Parse the games
	games := []*Game{}

	dir, err := os.Open(dirName)
	if err != nil {
		return nil, err
	}

	gameList, err := dir.Readdirnames(0)
	dir.Close()

	if err != nil {
		return nil, err
	}

	for _, game := range gameList {
		f, err := os.Stat(filepath.Join(dirName, game))

		// we only care about json files
		if err != nil || f.IsDir() || !strings.HasSuffix(game, ".json") {
			continue
		}

		gameFile, err := os.Open(filepath.Join(dirName, game))
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
