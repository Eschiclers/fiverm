package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var Force bool
var ResourcesFile string
var WorkingDirectory string

type Resource struct {
	Name       string `json:"name"`
	ZipballUrl string `json:"zipball_url"`
	Url        string `json:"url"`
	Version    string `json:"tag_name"`
	Folder     string `json:"folder"`
}

type Resources struct {
	Name      string     `json:"name"`
	Author    string     `json:"author"`
	Website   string     `json:"website"`
	Resources []Resource `json:"resources"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a resources.json file",
	Long:  `Create a resources.json file`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat(ResourcesFile)
		if os.IsNotExist(err) || Force {
			color.Green("Creating resources.json file")
			CreateResourcesJson(ResourcesFile)
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

	// Get the working directory
	WorkingDirectory, _ = os.Getwd()

	ResourcesFile = WorkingDirectory + string(os.PathSeparator) + "resources.json"
}

func CreateResourcesJson(file string) {

	_file, err := os.Create(file)
	if err != nil {
		color.Red("Error creating file \n%s", err)
		os.Exit(1)
	}

	var projectName, projectAuthor, projectWebsite string
	fmt.Print("Name of the project/server: ")
	fmt.Scanln(&projectName)
	fmt.Print("Author of the project/server: ")
	fmt.Scanln(&projectAuthor)
	fmt.Print("Website of the project/server: ")
	fmt.Scanln(&projectWebsite)

	project := Resources{
		Name:      projectName,
		Author:    projectAuthor,
		Website:   projectWebsite,
		Resources: []Resource{},
	}
	color.Green("Project registered")

	// Struct to json idented
	jsonData, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		color.Red("Error creating json file \n%s", err)
		os.Exit(1)
	}
	_file.Write(jsonData)

	defer _file.Close()
	color.Green("Created file %s", file)
}
