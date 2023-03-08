package tools

import "github.com/urfave/cli/v2"

var Cmd = &cli.Command{
	Name:            "tools",
	Aliases:         []string{"n"},
	Usage:           "工具集",
	Action:          Action,
	SkipFlagParsing: false,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "out_path",
			Aliases: []string{"o"},
			Usage:   "指定Go项目输出目录",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    "rename",
			Aliases: []string{"r"},
			Usage:   "执行需要重命名的目录路径",
			Value:   "",
		},
	},
}
