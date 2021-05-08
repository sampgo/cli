package resource

import (
	"io/ioutil"
	"sampgo-cli/notify"
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

func Exists(fileName string) (bool, error) {
	_, err := ioutil.ReadFile(fileName)

	if err == nil {
		// sampgo.toml (or fileName) already exists in the current directory.
		return true, nil
	}

	// sampgo.toml (or fileName) doesn't exists in the current directory.
	return false, err
}
