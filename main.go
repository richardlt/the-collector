package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/richardlt/the-collector/server"
)

func main() {
	app := cli.NewApp()
	app.Name = "the-collector"

	var databaseURI string
	var databaseName string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "database-uri, db-uri",
			Value:       "localhost:27017",
			Destination: &databaseURI,
		},
		cli.StringFlag{
			Name:        "database-name, db-name",
			Value:       "the-collector",
			Destination: &databaseName,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "start the app",
			Action: func(c *cli.Context) {
				server.Start(databaseURI, databaseName)
			},
		},
	}

	app.Run(os.Args)
}
