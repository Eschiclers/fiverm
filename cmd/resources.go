package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Global variables
var ResourcesFile string
var WorkingDirectory string
var TemporalFolder string

// Structs
type Resource struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Author     string `json:"author"`
	Repository string `json:"repository"`
	Url        string `json:"url"`
	Folder     string `json:"folder"`
	ZipballUrl string `json:"zipball_url"`
}

type Resources struct {
	Name      string     `json:"name"`
	Author    string     `json:"author"`
	Website   string     `json:"website"`
	Resources []Resource `json:"resources"`
}

func CreateResourcesFile() {
	_file, err := os.Create(ResourcesFile)
	if err != nil {
		color.Red("Error creating file \n%s", err)
		os.Exit(1)
	}

	var projectName, projectAuthor, projectWebsite string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Name of the project/server: ")
	projectName, _ = reader.ReadString('\n')
	projectName = strings.TrimSpace(projectName)
	fmt.Print("Author of the project/server: ")
	projectAuthor, _ = reader.ReadString('\n')
	projectAuthor = strings.TrimSpace(projectAuthor)
	fmt.Print("Website of the project/server: ")
	projectWebsite, _ = reader.ReadString('\n')
	projectWebsite = strings.TrimSpace(projectWebsite)

	project := Resources{
		Name:      projectName,
		Author:    projectAuthor,
		Website:   projectWebsite,
		Resources: []Resource{},
	}

	// Struct to json idented
	jsonData, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		color.Red("Error creating json file \n%s", err)
		os.Exit(1)
	}
	_file.Write(jsonData)
	defer _file.Close()

	color.Green("Project created")
}
