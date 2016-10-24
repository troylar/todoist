package main

import (
	"errors"
	"github.com/urfave/cli"
	"os"
)

var (
	default_config_path = os.Getenv("HOME") + "/.todoist.config.json"
	default_cache_path  = os.Getenv("HOME") + "/.todoist.cache.json"
	CommandFailed       = errors.New("Command Failed")
)

func main() {
	config, err := Setup(default_config_path)
	if err != nil {
		return
	}
	app := cli.NewApp()
	app.Name = "todoist"
	app.Usage = "Todoist cli client"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "Shows all tasks",
			Action: func(c *cli.Context) error {
				return List(config, c)
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add task",
			Action: func(c *cli.Context) error {
				return Add(config, c)
			},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "Complete task",
			Action: func(c *cli.Context) error {
				return Complete(config, c)
			},
		},
	}
	app.Run(os.Args)
}
