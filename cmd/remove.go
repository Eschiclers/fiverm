package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove [resource]",
	Short:   "Remove a resource",
	Long:    `Remove a resource`,
	Aliases: []string{"rm", "delete", "del", "uninstall"},
	Run: func(cmd *cobra.Command, args []string) {
		LoadResourcesFile()
		if len(args) == 0 {
			color.Red("Please specify a resource to remove.")
			os.Exit(1)
		}

		for _, arg := range args {
			if !ResourceInstalled(arg) {
				color.Red("The resource '" + arg + "' does not exist.")
				os.Exit(1)
			} else {
				RemoveResource(arg)
			}
		}

		SaveResourcesFile()
		color.Green("The resource(s) have been removed.")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
