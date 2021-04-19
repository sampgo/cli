package handler

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sampgo-cli/config"
	"sampgo-cli/notify"

	"github.com/urfave/cli/v2"
)

func defaultBuild(c config.Config) error {
	os.Setenv("CGO_CFLAGS_ALLOW", ".*")
	os.Setenv("CGO_LDFLAGS_ALLOW", ".*")

	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("GOARCH", "386")

	if runtime.GOOS == "windows" {
		// set GOOS to windows on windows.
		os.Setenv("GOOS", "windows")
	} else {
		// any other OS? set it to linux.
		os.Setenv("GOOS", "linux")
	}

	// For the time being, we will keep verbose mode persistent.
	cmd := fmt.Sprintf("go build -x -buildmode=c-shared -o %s %s", c.Package.Output, c.Package.Output)
	_, err := exec.Command(cmd).Output()

	if err != nil {
		return err
	}

	return nil
}

func sampctlBuild(c config.Config) error {
	err := defaultBuild(c)

	if err != nil {
		return err
	}

	return nil
}

func Build(ctx *cli.Context) error {
	fileName := "sampgo.toml"

	// Create a new instance of our parser.
	p := config.Parser{
		Dialect:  config.Toml,
		FileName: fileName,
	}

	// Parse our toml configuration file.
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
