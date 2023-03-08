package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/lemonlzy/flora/new"
	"github.com/lemonlzy/flora/tools"
)

var Version string

func main() {
	app := &cli.App{
		Name:    "flora",
		Usage:   "flora脚手架工具",
		Version: Version,
		Commands: []*cli.Command{
			new.Cmd,
			tools.Cmd,
			//add.Cmd,
			//update.Cmd,
			//upgrade.Cmd,
			//mysql.Cmd,
			//fix.Cmd,
		},
		Before: nil,
	}
	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
