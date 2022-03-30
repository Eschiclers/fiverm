package cmd

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var TemporalFolder string

type Response struct {
	Name       string
	Version    string `json:"tag_name"`
	ZipballUrl string `json:"zipball_url"`
}

var installCmd = &cobra.Command{
	Use:   "install [resources]",
	Short: "Install a resource",
	Long:  `Install a resource from github repository`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		_, err := os.Stat(ResourcesFile)
		if os.IsNotExist(err) {
			color.Red("The resources.json file does not exist")
			color.Yellow("Use fiverm init to create it")
			os.Exit(1)
		}
		_, err = os.Stat(WorkingDirectory + string(os.PathSeparator) + "resources")
		if os.IsNotExist(err) {
			color.Red("The resources folder does not exist")
			color.Yellow("Make soure you are in the right directory")
			os.Exit(1)
		}

		_, err = os.Stat(TemporalFolder)
		if os.IsNotExist(err) {
			os.Mkdir(TemporalFolder, 0755)
		}

		// TODO: Add support for versions with @version | example: fivemtools/ft_ui@0.1 | example: fivemtools/ft_ui@latest
		for i := 0; i < len(args); i++ {
			git_url := "https://api.github.com/repos/" + args[i] + "/releases/latest"
			// do request and save json
			resp, err := http.Get(git_url)
			if err != nil {
				color.Red("%s", err)
				os.Exit(1)
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				color.Red("Can not get the latest release of '" + args[i] + "'")
				color.Red("Error: " + resp.Status)
				color.Yellow("You may want to download the files from the master branch?")
				color.Yellow("Use the flag --master 'fiverm install " + args[i] + " --master'")
				os.Exit(1)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				color.Red("%s", err)
				os.Exit(1)
			}
			//Convert the body to type string
			sb := string(body)

			var resource Resource
			err = json.Unmarshal([]byte(sb), &resource)
			if err != nil {
				color.Red("%s", err)
				os.Exit(1)
			}
			resource.Name = strings.Split(args[i], "/")[1]

			// Found the resource
			color.Green("Found the resource '" + args[i] + "' (version: " + resource.Version + ")")

			color.Blue("Downloading the latest release")

			// Download the zip file
			err = downloadFile(resource.ZipballUrl, TemporalFolder+resource.Name+".zip")
			if err != nil {
				color.Red("Can not download the zip file")
				color.Red("Error: %s", err)
				os.Exit(1)
			}
			color.Green("Downloaded '" + args[i] + "'")

			// Unzip the file in the temporal folder
			color.Blue("Unzipping the zip file")
			// parent is the folder name where the resource will be unzipped
			parent, err := unzipSource(TemporalFolder+resource.Name+".zip", TemporalFolder+string(os.PathSeparator))
			if err != nil {
				color.Red("Can not unzip the zip file")
				color.Red("Error: %s", err)
				os.Exit(1)
			}
			color.Green("Unzipped '" + args[i] + "' " + parent)
		}
		defer os.RemoveAll(TemporalFolder)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().StringP("folder", "", "", "The folder to install the resource/s")
	installCmd.Flags().BoolP("master", "m", false, "Install the resource/s from master branch")

	TemporalFolder = os.TempDir() + string(os.PathSeparator) + "fiverm" + string(os.PathSeparator)
}

/*
  Download file from URL to destination path
  @param url string
  @param destination string
  @return error
*/
func downloadFile(url string, filepath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return err
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

/*
  Unzip file from the source to destination path
  @param source string
  @param destination string
  @return error
*/
func unzipSource(source, destination string) (string, error) {
	var parentFolder string
	// Open the zip file
	reader, err := zip.OpenReader(source)
	if err != nil {
		return "", err
	}
	defer reader.Close()

	// Iterate over zip files inside the archive and unzip each of them
	for i, f := range reader.File {
		if i == 0 {
			parentFolder = filepath.Dir(f.Name)
		}
		err := unzipFile(f, destination)
		if err != nil {
			return "", err
		}
	}

	return parentFolder, nil
}

func unzipFile(f *zip.File, destination string) error {
	// Check if file paths are not vulnerable to Zip Slip
	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	// Create directory tree
	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// Create a destination file for unzipped content
	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// Unzip the content of a file and copy it to the destination file
	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}
	return nil
}
