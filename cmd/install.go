package cmd

import (
	"encoding/json"
	"fmt"
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
	ZipballUrl string `json:"zipball_url"`
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a resource",
	Long:  `Install a resource`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat(ResourcesFile)
		if os.IsNotExist(err) {
			color.Red("The resources.json file does not exist")
			color.Yellow("Use fiverm init to create it")
			os.Exit(1)
		}

		_, err = os.Stat(TemporalFolder)
		if os.IsNotExist(err) {
			os.Mkdir(TemporalFolder, 0755)
		}

		// TODO: Add support for versions with @version | example: fivemtools/ft_ui@0.1 | example: fivemtools/ft_ui@latest
		for i := 0; i < len(args); i++ {
			/*
				resource := strings.Split(args[i], "@")
				var repository, version string
				repository = resource[0]
				if len(resource) > 1 {
					version = resource[1]
				} else {
					version = "latest"
				}

				fmt.Println("Resource " + repository)
				fmt.Println("Version " + version)
			*/

			git_url := "https://api.github.com/repos/" + args[i] + "/releases/latest"
			fmt.Println(git_url)
			// do request and save json
			resp, err := http.Get(git_url)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			//Convert the body to type string
			sb := string(body)

			var response Response
			err = json.Unmarshal([]byte(sb), &response)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Download the file
			err = DownloadFile(response.ZipballUrl, WorkingDirectory+string(os.PathSeparator)+"tmp"+string(os.PathSeparator)+strings.Split(args[i], "/")[1]+".zip")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println("Downloaded " + args[i] + ".zip")
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().StringP("folder", "", "", "The folder to install the resource/s")
	installCmd.Flags().BoolP("master", "m", false, "Install the resource/s from master branch")

	TemporalFolder = WorkingDirectory + string(os.PathSeparator) + "tmp"
}

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
