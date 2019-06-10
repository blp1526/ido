package main

import (
	"os"

	"github.com/urfave/cli"
)

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

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
