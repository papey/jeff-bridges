package internal

import (
	"errors"

	"github.com/milosgajdos83/tenus"

	"github.com/urfave/cli"
)

// HandleDelbr handles delete command
func HandleDelbr(c *cli.Context) error {
	// check argument number
	if c.NArg() != 1 {
		return errors.New("Name argument is missing for delete command")
	}

	// check if br exists
	br, err := tenus.BridgeFromName(c.Args().First())
	if err != nil {
		return err
	}

	// delete and check error
	return br.DeleteLink()

}
