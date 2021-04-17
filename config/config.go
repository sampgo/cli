package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/komkom/toml"
	"io/ioutil"
	"strings"
)

type CfgDialect int

const (
	PawnCfg CfgDialect = iota
	Toml
)

type Parser struct {
	Dialect  CfgDialect
	FileName string
}

func New(fileName string, dialect CfgDialect) Parser {
	return Parser{
		Dialect:  dialect,
		FileName: fileName,
	}
}

func (p Parser) Parse(Struct interface{}) (interface{}, error) {
	if strings.HasSuffix(p.FileName, ".cfg") || strings.HasSuffix(p.FileName, ".toml") {
		// This is just a really simple way of checking for files that already have an extension.
		// In future this will be replaced with regex. For now, this is OK!
		return nil, fmt.Errorf("file ext already inferred from cfg dialect")
	}

	var ext string

	// No ternary operators? Fine! :grin:
	if p.Dialect == PawnCfg {
		ext = ".cfg"
	} else if p.Dialect == Toml {
		ext = ".toml"
	} else {
		// No extension? kek
		ext = ""
	}

	data, err := ioutil.ReadFile(p.FileName + ext)
	if err != nil {
		return nil, fmt.Errorf("invalid file name")
	}

	// Our toml library leverages the default JSON decoder.
	decoder := json.NewDecoder(toml.New(bytes.NewBuffer(data)))

	return nil, nil
}
