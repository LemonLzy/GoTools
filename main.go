package main

import (
	"github.com/lemonlzy/flora/pretty"
	"github.com/lemonlzy/flora/rename"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var Version string

func main() {
	app := &cli.App{
		Name:    "flora",
		Usage:   "flora脚手架工具",
		Version: Version,
		Commands: []*cli.Command{
			//new.Cmd,
			rename.Cmd,
			pretty.Cmd,
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
