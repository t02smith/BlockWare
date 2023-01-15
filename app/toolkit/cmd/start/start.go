/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package start

import (
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a process",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

}
