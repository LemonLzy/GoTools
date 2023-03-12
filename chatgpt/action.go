package chat

import (
	"github.com/urfave/cli/v2"
)

func Action(cli *cli.Context) error {
	Run()
	return nil
}
