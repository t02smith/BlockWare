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

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// gameCmd represents the game command
var gameCmd = &cobra.Command{
	Use:   "game",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error creating game")
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

		_, err = titlePrompt.Run()
		if err != nil {
			fmt.Println("Error creating game")
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

		_, err = versionPrompt.Run()
		if err != nil {
			fmt.Println("Error creating game")
		}

		// root directory

	},
}

func init() {
	NewCmd.AddCommand(gameCmd)
}
