package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v2"
)

type GameType string

const (
	Classic GameType = "_classic_"
	Retail           = "_retail_"
)
const (
	InterfacePath = "Interface/AddOns"
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
	GameDir   string              `yaml:"gameDir"`
	GameType
}

func (c *Config) AddonDir() string {
	return fmt.Sprintf("%s/%s/%s", c.GameDir, c.GameType, InterfacePath)
}

func ReadFile(file string) Config {
	c := Config{}
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	unmarshalError := yaml.Unmarshal(b, &c)
	if unmarshalError != nil {
		log.Fatal(err.Error())
	}
	return c
}
