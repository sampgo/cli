package config

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
