package internal

import (
	"errors"
	"net"

	"github.com/milosgajdos83/tenus"

	"github.com/papey/jeff-bridges/pkg/networks"

	"github.com/urfave/cli"
)

// HandleAddbr handles add command
func HandleAddbr(c *cli.Context) error {
	// check argument number
	if c.NArg() < 2 {
		return errors.New("Add command is missing some argument(s)")
	}

	// is base network valid ?
	_, _, err := net.ParseCIDR(c.Args().Get(1))

	_, err = networks.GetAll()
	if err != nil {
		return err
	}

	// TODO:
	// - Check if next subnet is ok
	// - If ok, create bridge using name
	// - Assign IP

	return err
}

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
