package rename

import "github.com/urfave/cli/v2"

var Cmd = &cli.Command{
	Name:            "rename",
	Aliases:         []string{"r"},
	Usage:           "文件重命名",
	Action:          Action,
	SkipFlagParsing: false,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "path",
			Aliases: []string{"p"},
			Usage:   "指定需要替换文件名的目录",
			Value:   "",
		},
	},
}
