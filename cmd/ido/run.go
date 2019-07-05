package main

import (
	"os"

	"github.com/blp1526/ido"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:      "run",
	Usage:     "Runs a container",
	ArgsUsage: "[image] [command]",
	Action: func(c *cli.Context) (err error) {
		if len(c.Args()) != 2 {
			err := cli.ShowCommandHelp(c, "run")
			if err != nil {
				return cli.NewExitError(err, exitCodeNG)
			}

			return nil
		}

		image := c.Args()[0]
		command := c.Args()[1]

		dir, err := ido.Create(image)
		if err != nil {
			return cli.NewExitError(err, exitCodeNG)
		}
		defer os.RemoveAll(dir) // nolint: errcheck

		err = ido.Run(dir, command)
		if err != nil {
			return cli.NewExitError(err, exitCodeNG)
		}

		return nil
	},
}
