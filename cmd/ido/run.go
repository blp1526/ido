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
	Flags: []cli.Flag{
		cli.StringSliceFlag{
			Name:  "volume, v",
			Usage: "Create a bind mount `HOST-DIR:CONTAINER-DIR`",
		},
	},
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
		volumes := c.StringSlice("volume")

		dir, err := ido.Create(image)
		if err != nil {
			return cli.NewExitError(err, exitCodeNG)
		}
		defer os.RemoveAll(dir) // nolint: errcheck

		err = ido.Run(dir, command, volumes)
		if err != nil {
			return cli.NewExitError(err, exitCodeNG)
		}

		return nil
	},
}
