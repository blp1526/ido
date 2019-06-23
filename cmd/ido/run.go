package main

import (
	"fmt"

	"github.com/blp1526/ido"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:      "run",
	Usage:     "runs a container",
	ArgsUsage: "[image]",
	Action: func(c *cli.Context) (err error) {
		image := c.Args().First()
		if image == "" {
			err := cli.ShowCommandHelp(c, "run")
			if err != nil {
				return cli.NewExitError(err, exitCodeNG)
			}

			return nil
		}

		result, err := ido.Create(image)
		if err != nil {
			return cli.NewExitError(err, exitCodeNG)
		}

		fmt.Printf("%s\n", result)
		return nil
	},
}
