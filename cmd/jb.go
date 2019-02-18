package main

import (
	"fmt"
	"log"
	"os"

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
				fmt.Println("[TODO] Add Bridge")
				return nil
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "delete a bridge",
			Action: func(c *cli.Context) error {
				fmt.Println("[TODO] Del Bridge")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}