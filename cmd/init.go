package cmd

import (
	"os"

	"github.com/Eschiclers/fiverm/pkg/workingdirectory"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a fiverm.json file",
	Long:  `Create a fiverm.json file`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(workingdirectory.GetWorkingDirectory() + string(os.PathSeparator) + "resources.json"); os.IsNotExist(err) {
			color.Green("Creating resources.json file")
		} else {
			color.Red("resource.json file already exists")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
