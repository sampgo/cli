package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sampgo-cli/notify"
	"sampgo-cli/resource"
	"sampgo-cli/sampctl"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

var questions = []*survey.Question{
	{
		Name:     "username",
		Prompt:   &survey.Input{Message: "What is your username?"},
		Validate: survey.Required,
	},
	{
		Name:     "repo",
		Prompt:   &survey.Input{Message: "What is your repo called?"},
		Validate: survey.Required,
	},
	{
		Name: "gomode",
		Prompt: &survey.Input{Message: "Enter your gomode name:", Default: func() string {
			path, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}

			cwd := strings.Split(path, string(os.PathSeparator)) // current working dir
			dir := cwd[len(cwd)-1]

			return dir
		}()},
		Validate: survey.Required,
	},
	{
		Name: "input",
		Prompt: &survey.Input{
			Message: "Enter your gomode entrypoint file name:",
			Suggest: func(toComplete string) []string {
				files, _ := filepath.Glob(toComplete + "*.go")
				return files
			},
		},
		Validate: survey.Required,
	},
	{
		Name: "output",
		Prompt: &survey.Input{
			Message: "Enter your desired gomode output path:",
			Default: func() string {
				path, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}

				cwd := strings.Split(path, string(os.PathSeparator)) // current working dir
				dir := cwd[len(cwd)-1]

				var ext string

				if runtime.GOOS == "windows" {
					// Windows.
					ext = "dll"
				} else {
					// Unix-based system.
					ext = "so"
				}

				return fmt.Sprintf("../plugins/%s.%s", dir, ext)
			}(),
		},
		Validate: survey.Required,
	},
}

func Init(c *cli.Context) error {
	fileName := "sampgo.toml"

	_, err := ioutil.ReadFile(fileName)
	if err == nil {
		// sampgo.toml (or fileName) already exists in the current directory.
		notify.Error("A sampgo package already exists in your directory.")
		return err
	}

	sampctlFound := sampctl.Installed()

	if sampctlFound {
		// sampctl is available on this system.
		notify.Info("sampctl found, defaulting to sampctl support.")
	} else {
		// sampctl is not available on this system.
		notify.Info("sampctl not found, sampctl support will be toggled off.")
		notify.Warning("It is advised that you install sampctl, as it will enhance your sampgo experience.")
	}

	// start structuring our "survey" per-se

	answers := struct {
		Username string // Author username.
		Repo     string // gomode repo.
		Gomode   string // gomode name.
		Input    string // gomode entrypoint file.
		Output   string // gomode output file.
	}{}

	// fucking hell...
	err = survey.Ask(questions, &answers)
	if err != nil {
		// One question probably wasn't answered.
		notify.Error(err.Error())
		return nil
	}

	conf := resource.Config{}

	conf.Global.Sampctl = sampctlFound

	conf.Author.User = answers.Username
	conf.Author.Repo = answers.Repo

	conf.Package.Name = answers.Gomode
	conf.Package.Input = answers.Input
	conf.Package.Output = answers.Output

	err = resource.WriteToml(fileName, conf)
	if err != nil {
		notify.Error("sampgo configuration could not be written to.")
		return nil
	}

	return nil
}
