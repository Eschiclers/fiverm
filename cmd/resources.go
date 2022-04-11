package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Global variables
var ResourcesFile string
var WorkingDirectory string
var TemporalFolder string
var Project Resources

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

func LoadResourcesFile() {
	if !ResourcesFileExists() {
		color.Red("The resources.json file does not exist")
		color.Yellow("Use fiverm init to create it")
		os.Exit(1)
	}
	file, _ := ioutil.ReadFile(ResourcesFile)
	json.Unmarshal(file, &Project)
}

func SaveResourcesFile() {
	// Struct to json idented
	jsonData, err := json.MarshalIndent(Project, "", "  ")
	if err != nil {
		color.Red("Error creating json file \n%s", err)
		os.Exit(1)
	}

	_file, err := os.Create(ResourcesFile)
	if err != nil {
		color.Red("Error creating file \n%s", err)
		os.Exit(1)
	}
	_file.Write(jsonData)
	defer _file.Close()
}

func ResourcesFileExists() bool {
	_, err := os.Stat(ResourcesFile)
	return !os.IsNotExist(err)
}

func ResourcesFolderExists() bool {
	_, err := os.Stat(WorkingDirectory + string(os.PathSeparator) + "resources")
	return !os.IsNotExist(err)
}

func ResourceInstalled(name string) bool {
	// If the string contains a slash, get the last part of the string
	if strings.Contains(name, "/") {
		name = strings.Split(name, "/")[len(strings.Split(name, "/"))-1]
	}

	for _, resource := range Project.Resources {
		if resource.Name == name {
			return true
		}
	}
	return false
}

func RemoveResource(name string) {
	// If the string contains a slash, get the last part of the string
	if strings.Contains(name, "/") {
		name = strings.Split(name, "/")[len(strings.Split(name, "/"))-1]
	}

	for i, resource := range Project.Resources {
		if resource.Name == name {
			color.Blue("Removing resource %s", resource.Name)
			var resourceFolder string

			if resource.Folder != "" {
				resourceFolder = resource.Folder + string(os.PathSeparator) + resource.Name
			} else {
				resourceFolder = resource.Name
			}

			err := os.RemoveAll(WorkingDirectory + string(os.PathSeparator) + "resources" + string(os.PathSeparator) + resourceFolder)
			if err != nil {
				color.Red("Error removing resource %s", resource.Name)
				color.Red(err.Error())
				os.Exit(1)
			}
			Project.Resources = append(Project.Resources[:i], Project.Resources[i+1:]...)
			color.Green("Resource %s removed", resource.Name)
			break
		}
	}
}
