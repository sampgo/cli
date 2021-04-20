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

func Exists(fileName string) bool {
	_, err := ioutil.ReadFile(fileName)
	if err == nil {
		// sampgo.toml (or fileName) already exists in the current directory.
		notify.Error("A sampgo package already exists in your directory.")
		return false
	}

	return true
}
