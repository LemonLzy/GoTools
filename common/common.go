package common

import (
	"errors"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

func CheckArgs(cli *cli.Context, command string) error {
	args := cli.Args()
	if args.Len() < 0 {
		color.Red.Printf("%s Command Line new execution error, please use 'flora %s -h for details\n'", module, command)
		return errors.New("lack arguments")
	}
	return nil
}
