package config

import (
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v2"
)

type Addon struct {
	Source string `yaml:"source,omitempty"`
	Id     string `yaml:"id,omitempty"`
}

type Provider struct {
	ClassicAddons map[string]Addon `yaml:"classic"`
	RetailAddons  map[string]Addon `yaml:"retail"`
}

type Config struct {
	Providers map[string]Provider `yaml:"providers"`
}

func FromFile(file string) (*Config, error) {
	c := new(Config)
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return c, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return c, err
	}
	unmarshalError := yaml.Unmarshal(b, c)
	if unmarshalError != nil {
		return c, err
	}
	return c, nil
}
