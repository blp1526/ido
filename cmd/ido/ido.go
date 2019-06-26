package main

import (
	"fmt"
	"os"

	"github.com/blp1526/ido"
	"github.com/urfave/cli"
)

const exitCodeNG = 1

func main() {
	app := cli.NewApp()
	app.Name = "ido"
	app.Version = ido.Version()
	app.Usage = ""
	app.Description = fmt.Sprintf("REVISION: %s", ido.Revision())
	app.Authors = []cli.Author{
		{
			Name:  "Shingo Kawamura",
			Email: "blp1526@gmail.com",
		},
	}
	app.Copyright = "(c) 2019 Shingo Kawamura"

	app.Commands = []cli.Command{
		createCommand,
		runCommand,
		attachCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
