package main

import (
	"github.com/blp1526/ido"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:      "run",
	Usage:     "runs a container",
	ArgsUsage: "[command]",
	Action: func(c *cli.Context) (err error) {
		cmd := c.Args().First()
		if cmd == "" {
			err := cli.ShowCommandHelp(c, "run")
			if err != nil {
				return cli.NewExitError(err, exitCodeNG)
			}

			return nil
		}

		err = ido.Run(cmd)
		if err != nil {
			return cli.NewExitError(err, exitCodeNG)
		}

		return nil
	},
}
