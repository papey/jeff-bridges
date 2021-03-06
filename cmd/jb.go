package main

import (
	"log"
	"os"

	internal "github.com/papey/jeff-bridges/internal/app"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "jeff-bridges"
	app.Usage = "Multiply bridges for fun and profit"

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a bridge",
			Action: func(c *cli.Context) error {
				return internal.HandleAddbr(c)
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "delete a bridge",
			Action: func(c *cli.Context) error {
				return internal.HandleDelbr(c)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
