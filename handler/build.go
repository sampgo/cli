package handler

import (
	"fmt"
	"os"
	"os/exec"
	"sampgo-cli/env"
	"sampgo-cli/notify"
	"sampgo-cli/resource"

	"github.com/urfave/cli/v2"
)

func defaultBuild(c resource.Config, v bool) error {
	env.Set()

	var args []string

	// For the time being, we will keep verbose mode persistent.
	if v {
		// verbose mode enabled
		args = []string{"bash", "-c", "go build", "-x", "-buildmode=c-shared", "-o", c.Package.Output, c.Package.Input}

		notify.Info(fmt.Sprintf("using %s as entrypoint file", c.Package.Input))
		notify.Info(fmt.Sprintf("setting output to %s", c.Package.Output))
	} else {
		args = []string{"bash", "-c", "go build", "-buildmode=c-shared", "-o", c.Package.Output, c.Package.Input}
	}

	path, _ := exec.LookPath("bash")
	cmd := &exec.Cmd{
		Path:   path,
		Args:   args,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	err := cmd.Run()

	if err != nil {
		env.Unset()
		fmt.Println(cmd.String())

		return err
	}

	env.Unset()

	return nil
}

func sampctlBuild(c resource.Config, v bool) error {
	err := defaultBuild(c, v)

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
		sampctlBuild(c, ctx.Bool("verbose"))
	} else {
		// Build without the sampctl stuff.
		// TODO: handle error.
		defaultBuild(c, ctx.Bool("verbose"))
	}

	return nil
}
