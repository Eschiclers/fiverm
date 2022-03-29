package cmd

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
	Long:  `Install a resource`,
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
			var zipFile string

			git_url := "https://api.github.com/repos/" + args[i] + "/releases/latest"
			// do request and save json
			resp, err := http.Get(git_url)
			if err != nil {
				color.Red("%s", err)
				os.Exit(1)
			}
			if resp.StatusCode != 200 {
				color.Red("Can not get the latest release of '" + args[i] + "'")
				color.Red("Error: " + resp.Status)
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

			// Download the file
			zipFile = TemporalFolder + string(os.PathSeparator) + strings.Split(args[i], "/")[1] + ".zip"
			err = DownloadFile(resource.ZipballUrl, zipFile)
			if err != nil {
				color.Red("%s", err)
				os.Exit(1)
			}

			color.Green("Downloaded '" + args[i] + "'")
		}
		defer os.RemoveAll(TemporalFolder)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().StringP("folder", "", "", "The folder to install the resource/s")
	installCmd.Flags().BoolP("master", "m", false, "Install the resource/s from master branch")

	TemporalFolder = os.TempDir() + string(os.PathSeparator) + "fiverm"
}

/*
  Download file from URL to destination path
  @param url string
  @param destination string
  @return error
*/
func DownloadFile(url string, filepath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
