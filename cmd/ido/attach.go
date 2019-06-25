package main

import (
	"github.com/blp1526/ido"
	"github.com/urfave/cli"
)

var attachCommand = cli.Command{
	Name:      "attach",
	Usage:     "attaches to a container",
	ArgsUsage: "[pid]",
	Action: func(c *cli.Context) (err error) {
		pid := c.Args().First()
		if pid == "" {
			err := cli.ShowCommandHelp(c, "attach")
			if err != nil {
				return cli.NewExitError(err, exitCodeNG)
			}

			return nil
		}

		err = ido.Attach(pid)
		if err != nil {
			return cli.NewExitError(err, exitCodeNG)
		}

		return nil
	},
}
