package main

import (
	"fmt"

	"github.com/blp1526/ido"
	"github.com/urfave/cli"
)

var createCommand = cli.Command{
	Name:      "create",
	Usage:     "creates a container",
	ArgsUsage: "[image]",
	Action: func(c *cli.Context) (err error) {
		image := c.Args().First()
		if image == "" {
			err := cli.ShowCommandHelp(c, "create")
			if err != nil {
				return cli.NewExitError(err, exitCodeNG)
			}

			return nil
		}

		tempDir, err := ido.Create(image)
		if err != nil {
			return cli.NewExitError(err, exitCodeNG)
		}

		fmt.Println(tempDir)
		return nil
	},
}
