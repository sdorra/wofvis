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
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "use-openpgp-api",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "serve",
			Usage:  "starts a webserver and shows the web of trust graph",
			Action: pkg.Serve,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "addr",
					Usage:  "listen address for the webserver",
					Value:  "127.0.0.1:8080",
					EnvVar: "WOFVIS_ADDR",
				},
			},
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
