package handler

import (
	"fmt"
	"os/exec"
	"sampgo-cli/env"
	"sampgo-cli/notify"
	"sampgo-cli/resource"

	"github.com/urfave/cli/v2"
)

func defaultBuild(c resource.Config) error {
	env.Set()

	// For the time being, we will keep verbose mode persistent.
	cmd := fmt.Sprintf("go build -x -buildmode=c-shared -o %s %s", c.Package.Output, c.Package.Input)
	_, err := exec.Command(cmd).Output()

	if err != nil {
		env.Unset()
		return err
	}

	return nil
}

func sampctlBuild(c resource.Config) error {
	err := defaultBuild(c)

	if err != nil {
		return err
	}

	return nil
}

func Build(ctx *cli.Context) error {
	fileName := "sampgo.toml"

	// Create a new instance of our parser.
	p := resource.Parser{
		Dialect:  resource.Toml,
		FileName: fileName,
	}

	// Parse our toml resourceuration file.
	c, err := p.ParseToml()
	if err != nil {
		notify.Error("sampgo.toml was unable to be parsed!")
		return nil
	}

	if c.Global.Sampctl {
		// begin the sampctl related stuff.
		// TODO: handle error.
		sampctlBuild(c)
	} else {
		// Build without the sampctl stuff.
		// TODO: handle error.
		defaultBuild(c)
	}

	return nil
}
