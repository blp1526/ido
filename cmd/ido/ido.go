package main

import (
	"fmt"
	"os"

	"github.com/blp1526/ido"
	"github.com/urfave/cli/v2"
)

const exitCodeNG = 1

func main() {
	app := &cli.App{
		Name:        "ido",
		Version:     ido.Version(),
		Usage:       "",
		Description: fmt.Sprintf("REVISION: %s", ido.Revision()),
		Authors: []*cli.Author{
			{
				Name:  "Shingo Kawamura",
				Email: "blp1526@gmail.com",
			},
		},
		Copyright: "(c) 2019 Shingo Kawamura",
		Commands: []*cli.Command{
			createCommand,
			runCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
