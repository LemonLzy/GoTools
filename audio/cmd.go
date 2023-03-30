package audio

import "github.com/urfave/cli/v2"

var Cmd = &cli.Command{
	Name:            "audio",
	Aliases:         []string{"a"},
	Usage:           "音频转换",
	Action:          Action,
	SkipFlagParsing: false,
}
