/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove [resource]",
	Short:   "Remove a resource",
	Long:    `Remove a resource`,
	Aliases: []string{"rm", "delete", "del", "uninstall"},
	Run: func(cmd *cobra.Command, args []string) {
		LoadResourcesFile()
		color.Green(Project.Name)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
