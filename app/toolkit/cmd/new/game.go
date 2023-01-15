/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package new

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/t02smith/part-iii-project/toolkit/lib/games"
)

// gameCmd represents the game command
var gameCmd = &cobra.Command{
	Use:   "game",
	Short: "Create a new game",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error creating game: %s\n", err)
		}

		// game title
		titlePrompt := promptui.Prompt{
			Label:   "Game title",
			Default: filepath.Base(cwd),
			Validate: func(s string) error {
				if len(s) == 0 {
					return errors.New("too short")
				}
				return nil
			},
		}

		title, err := titlePrompt.Run()
		if err != nil {
			fmt.Printf("Error creating game: %s\n", err)
			return
		}

		// game version
		versionPrompt := promptui.Prompt{
			Label:   "Version",
			Default: "1.0.0",
			Validate: func(s string) error {
				if matches, err := regexp.MatchString("^(\\d+\\.)*\\d+$", s); !matches || err != nil {
					return errors.New("invalid version")
				}
				return nil
			},
		}

		version, err := versionPrompt.Run()
		if err != nil {
			fmt.Printf("Error creating game: %s\n", err)
			return
		}

		// developer

		devPrompt := promptui.Prompt{
			Label: "Your domain",
		}

		dev, err := devPrompt.Run()
		if err != nil {
			fmt.Printf("Error creating game: %s\n", err)
			return
		}

		// release date

		releasePrompt := promptui.Prompt{
			Label: "Release date (yyyy-mm-dd)",
			Validate: func(s string) error {
				m, err := regexp.MatchString("\\d{4}-\\d{2}-\\d{2}", s)
				if err != nil {
					return err
				}

				if !m {
					return errors.New("invalid date")
				}
				return nil
			},
		}

		release, err := releasePrompt.Run()
		if err != nil {
			fmt.Printf("Error creating game: %s\n", err)
			return
		}

		// root directory

		rootDirPrompt := promptui.Prompt{
			Label:   "Root directory of Game",
			Default: ".",
			Validate: func(s string) error {

				f, err := os.Stat(s)
				if err != nil {
					return err
				}

				if !f.IsDir() {
					return errors.New("not a dir")
				}

				return nil
			},
		}

		rootDir, err := rootDirPrompt.Run()
		if err != nil {
			fmt.Printf("Error creating game: %s\n", err)
			return
		}

		// shard size

		shardSizePrompt := promptui.Prompt{
			Label:   "Shard size (in bytes)",
			Default: "16384",
			Validate: func(s string) error {
				_, err := strconv.ParseUint(s, 10, 64)
				return err
			},
		}

		shardSize, err := shardSizePrompt.Run()
		if err != nil {
			fmt.Printf("Error creating game: %s\n", err)
			return
		}

		// create the game

		shardSizeUint, _ := strconv.ParseUint(shardSize, 10, 64)
		releaseStr, err := time.Parse("2006-01-02", release)
		if err != nil {
			fmt.Printf("Error creating game: %s\n", err)
			return
		}

		game, err := games.CreateGame(title, version, releaseStr.String(), dev, rootDir, uint(shardSizeUint))
		if err != nil {
			fmt.Printf("Error creating game: %s\n", err)
			return
		}

		// write game to file
		err = games.OutputToFile(game)
		if err != nil {
			fmt.Printf("Error writing game to file: %s\n", err)
		}
	},
}

func init() {
	NewCmd.AddCommand(gameCmd)
}
