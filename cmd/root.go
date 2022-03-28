package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fiverm",
	Short: "Fiverm is a CLI for managing your fivem server resources",
	Long: `Fiverm is a CLI for managing your fivem server resources.
This application is a tool to generate the needed files
to manage your resources. It allows you to add, update,
or remove resources among other functions.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
