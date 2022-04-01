package cmd

var ResourcesFile string

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
