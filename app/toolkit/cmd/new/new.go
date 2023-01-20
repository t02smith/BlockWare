/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package new

import (
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create something new",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

}
