package curseforge

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	baseURL    = "https://addons-ecs.forgesvc.net/api/v2/"
	categoryID = "0"
	gameID     = "1" // World of warcraft
	classic    = "wow_classic"
	retail     = "wow_retail"
)

// Match response from fingerprint search API
type Match struct {
	ID                 int    `json:"id"`
	File               File   `json:"file"`
	LatestFiles        []File `json:"latestFiles"`
	PackageFingerPrint uint32 `json:"packageFingerprint"`
	GameVersion        string `json:"gameVersion,omitempty"`
	ProjectID          int    `json:"projectId,omitempty`
}

type File struct {
	ID                int      `json:"id"`
	DownloadURL       string   `json:"downloadUrl"`
	ReleaseType       int      `json:"releaseType"`
	FileStatus        int      `json:"fileStatus"`
	IsAlternate       bool     `json:"isAlternate"`
	GameVersionFlavor string   `json:"gameVersionFlavor,omitempty"`
	Modules           []Module `json:"modules,omitempty"`
}

type Module struct {
	FolderName  string `json:"foldername"`
	FingerPrint uint32 `json:"fingerprint"`
	Type        int    `json:"type"`
}

type FingerPrintSearchResult struct {
	IsCacheBuilt   bool    `json:"isCacheBuilt"`
	ExactMatches   []Match `json:"exactMatches"`
	PartialMatches []Match `json:"partialMatches"`
}

type Addon struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	GameID      int    `json:"gameId"`
	Summary     string `json:"summary"`
	GameVersion string `json:"gameVersion"`
	LatestFiles []File `json:"latestFiles"`
}

type Provider struct {
}

func NewProvider() *Provider {
	return &Provider{}
}

func get(url string) ([]byte, error) {
	// TODO: Implement context
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	return body, nil
}

func post(url string, d []byte) ([]byte, error) {
	// TODO: Implement context

	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(d))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return body, nil
}

func Unpack(dir string) {

}
func Download() {

}

func (p *Provider) Update() {

}
