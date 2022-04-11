package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var Force bool

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a resources.json file",
	Long:  `Create a resources.json file`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat(ResourcesFile)
		if os.IsNotExist(err) || Force {
			color.Green("Creating resources.json file")
			CreateResourcesFile()
		} else {
			color.Red("The resource.json file already exists")
			color.Yellow("Use fiverm init -f to force overwriting")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&Force, "force", "f", false, "Force overwriting of the resources.json file")
}
