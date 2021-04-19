package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	toml2 "github.com/BurntSushi/toml"
	"github.com/komkom/toml"
)

type Global struct {
	Sampctl bool `json:"sampctl" toml:"sampctl"`
}

type Author struct {
	User string `json:"user" toml:"user"`
	Repo string `json:"repo" toml:"user"`
}

type Package struct {
	Name   string `json:"name" toml:"name"`
	Input  string `json:"input" toml:"input"`
	Output string `json:"output" toml:"output"`
}

// Config contains the basic TOML structure acceptable by sampgo.
type Config struct {
	Global  Global  `json:"global" toml:"global"`
	Author  Author  `json:"author" toml:"author"`
	Package Package `json:"package" toml:"package"`
}

// WriteToml allows you to write a toml file.
func WriteToml(fileName string, config Config) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	if err := toml2.NewEncoder(f).Encode(config); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
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
