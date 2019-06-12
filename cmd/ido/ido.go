package main

import (
	"fmt"
	"os"

	"github.com/blp1526/ido"
	"github.com/urfave/cli"
)

const exitCodeNG = 1

var version string

func main() {
	app := cli.NewApp()
	app.Name = "ido"
	app.Version = version
	app.Usage = ""
	app.Authors = []cli.Author{
		{
			Name:  "Shingo Kawamura",
			Email: "blp1526@gmail.com",
		},
	}
	app.Copyright = "(c) 2019 Shingo Kawamura"

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

	app.Commands = []cli.Command{
		runCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
