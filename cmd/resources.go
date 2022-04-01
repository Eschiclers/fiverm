package cmd

var ResourcesFile string

type Resource struct {
	Name       string `json:"name"`
	ZipballUrl string `json:"zipball_url"`
	Url        string `json:"url"`
	Version    string `json:"version"`
	Folder     string `json:"folder"`
}

type Resources struct {
	Name      string     `json:"name"`
	Author    string     `json:"author"`
	Website   string     `json:"website"`
	Resources []Resource `json:"resources"`
}
