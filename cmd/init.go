package cmd

import (
	"os"

	"github.com/Eschiclers/fiverm/pkg/workingdirectory"
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
		_, err := os.Stat(workingdirectory.GetWorkingDirectory() + string(os.PathSeparator) + "resources.json")

		if os.IsNotExist(err) || Force {
			color.Green("Creating resources.json file")
		} else {
			color.Red("The resource.json file already exists")
			color.Yellow("Use -f to force overwriting")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.PersistentFlags().BoolVarP(&Force, "force", "f", false, "Force to overwrite existing file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
