package chat

import "github.com/urfave/cli/v2"

var Cmd = &cli.Command{
	Name:            "chat",
	Aliases:         []string{"c"},
	Usage:           "工具集",
	Action:          Action,
	SkipFlagParsing: false,
}
