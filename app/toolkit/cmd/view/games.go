/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package view

import (
	"fmt"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/t02smith/part-iii-project/toolkit/lib/games"
)

// gamesCmd represents the games command
var gamesCmd = &cobra.Command{
	Use:   "games",
	Short: "Print out a table of your games",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		gameLs, err := games.LoadGames(viper.GetString("meta.directory"))
		if err != nil {
			log.Printf("Error loading games: %s\n", err)
			return
		}

		OutputGamesTable(gameLs)
	},
}

func init() {
	ViewCmd.AddCommand(gamesCmd)
}

func OutputGamesTable(gameLs []*games.Game) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Title", "Version", "Release"})

	for i, g := range gameLs {
		t.AppendRow(table.Row{fmt.Sprint(i + 1), g.Title, g.Version, g.ReleaseDate})
	}

	t.Render()
}
