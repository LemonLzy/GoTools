package pretty

import "github.com/urfave/cli/v2"

var Cmd = &cli.Command{
	Name:            "pretty",
	Aliases:         []string{"n"},
	Usage:           "工具集",
	Action:          Action,
	SkipFlagParsing: false,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "json",
			Aliases:     []string{"j"},
			Usage:       "未格式化的json字符串",
			Value:       "",
			Destination: &jsonInfo.json,
		},
	},
}
