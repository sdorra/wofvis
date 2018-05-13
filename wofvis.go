package main

import (
	"os"

	"github.com/sdorra/wofvis/pkg"
	"github.com/urfave/cli"
)

var Version = "0.0.0"

func main() {
	app := cli.NewApp()
	app.Name = "wofvis"
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:   "serve",
			Usage:  "starts a webserver and shows the usage graph",
			Action: pkg.Serve,
		},
		{
			Name:   "json",
			Usage:  "prints the web of trust as json to stdout",
			Action: pkg.PrintJSON,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
