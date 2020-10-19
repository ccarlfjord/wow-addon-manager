package curseforge

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseURL    = "https://addons-ecs.forgesvc.net/api/v2/addon/"
	categoryID = "0"
	gameID     = "1" // World of warcraft
	classic    = "wow_classic"
	retail     = "wow_retail"
)

type Provider struct {
}

type LatestFiles struct {
	ID                int    `json:"id"`
	DownloadURL       string `json:"downloadUrl"`
	ReleaseType       int    `json:"releaseType"`
	FileStatus        int    `json:"fileStatus"`
	IsAlternate       bool   `json:"isAlternate"`
	GameVersionFlavor string `json:"gameVersionFlavor"`
}

type SearchResult struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	GameID      int           `json:"gameId"`
	Summary     string        `json:"summary"`
	GameVersion string        `json:"gameVersion"`
	LatestFiles []LatestFiles `json:"latestFiles"`
}

func NewProvider() *Provider {
	return &Provider{}
}

// Search searches for addon on curseforge
func Search(name, version string) ([]SearchResult, error) {
	u, err := url.Parse(baseURL)
	u.Path += "search"
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Add("categoryID", categoryID)
	q.Add("gameID", gameID)
	q.Add("gameVersion", version)
	q.Add("searchFilter", name)
	q.Add("sectionID", "0")
	q.Add("sort", "0")
	u.RawQuery = q.Encode()
	fmt.Println(u.String())

	resp, err := get(u.String())
	if err != nil {
		return nil, err
	}
	s := make([]SearchResult, 0)
	if err := json.Unmarshal(resp, &s); err != nil {
		return s, err
	}
	return s, nil
}

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}
func (p *Provider) Install(dir string) {

}
func (p *Provider) Download() {

}

func (p *Provider) Update() {

}
