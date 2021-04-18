package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/komkom/toml"
	"io/ioutil"
)

// Config contains the basic TOML structure acceptable by sampgo.
type Config struct {
	Global struct {
		Sampctl bool `json:"sampctl"`
	} `json:"global"`

	Author struct {
		User string `json:"user"`
		Repo string `json:"repo"`
	} `json:"author"`

	Package struct {
		Name   string `json:"name"`
		Input  string `json:"input"`
		Output string `json:"output"`
	} `json:"package"`
}

// ParseToml parses the toml file specified by the user.
func (p Parser) ParseToml() (Config, error) {
	if p.Dialect != Toml {
		return Config{}, fmt.Errorf("instance of parser is not using toml")
	}

	data, err := ioutil.ReadFile(p.FileName)
	if err != nil {
		return Config{}, fmt.Errorf("invalid file name")
	}

	// Our toml library leverages the default JSON decoder.
	decoder := json.NewDecoder(toml.New(bytes.NewBuffer(data)))

	var config Config

	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
