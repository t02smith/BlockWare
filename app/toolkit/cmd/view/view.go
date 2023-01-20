/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package view

import (
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "Print some data out",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
